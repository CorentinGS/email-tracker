package postgres

import "time"

type Option func(*Postgres)

// WithMaxPoolSize sets the maximum number of connections in the pool.
func WithMaxPoolSize(maxPoolSize int) Option {
	return func(p *Postgres) {
		p.maxPoolSize = maxPoolSize
	}
}

// WithConnAttempts sets the number of connection attempts.
func WithConnAttempts(connAttempts int) Option {
	return func(p *Postgres) {
		p.connAttempts = connAttempts
	}
}

// WithConnTimeout sets the connection timeout.
func WithConnTimeout(connTimeout time.Duration) Option {
	return func(p *Postgres) {
		p.connTimeout = connTimeout
	}
}
