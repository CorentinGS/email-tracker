package postgres

import (
	"context"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	Pool         *pgxpool.Pool
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
}

var (
	postgresInstance *Postgres //nolint:gochecknoglobals //Singleton
	connOnce         sync.Once //nolint:gochecknoglobals //Singleton
)

func New(url string, opts ...Option) error {
	p := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	for _, opt := range opts {
		opt(p)
	}

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	// config.MaxConns = int32(p.maxPoolSize)
	config.ConnConfig.ConnectTimeout = p.connTimeout

	for i := 0; i < p.connAttempts; i++ {
		p.Pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err == nil {
			break
		}
	}

	if err != nil {
		return err
	}

	postgresInstance = p

	// test connection
	_, err = p.Pool.Acquire(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() {
	p.Pool.Close()
}

func GetPool() *pgxpool.Pool {
	return postgresInstance.Pool
}
