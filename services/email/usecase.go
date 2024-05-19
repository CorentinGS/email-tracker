package email

import (
	"context"

	db "github.com/corentings/email-tracker/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UseCase struct {
	q *db.Queries
}

func NewUseCase(dbConn *pgxpool.Pool) IUseCase {
	q := db.New(dbConn)

	return &UseCase{q: q}
}

func (uc *UseCase) GetEmail(ctx context.Context, imageUUID uuid.UUID) (db.Email, error) {
	return uc.q.GetEmailByUUID(ctx, imageUUID)
}

func (uc *UseCase) AddTracking(ctx context.Context, email db.Email, ip string) error {
	_, err := uc.q.InsertTracker(ctx, db.InsertTrackerParams{
		EmailUuid: email.Uuid,
		IpAddress: &ip,
	})

	return err
}

func (uc *UseCase) CreateEmail(ctx context.Context, recipient, subject string) (db.Email, error) {
	return uc.q.InsertEmail(ctx, db.InsertEmailParams{
		Recipient: recipient,
		Subject:   subject,
	})
}

func (uc *UseCase) GetEmailsWithPagination(ctx context.Context, limit int, offset int) ([]db.Email, error) {
	return uc.q.GetEmailsWithPagination(ctx, db.GetEmailsWithPaginationParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
}

func (uc *UseCase) GetTrackersWithPagination(ctx context.Context, limit int, offset int) ([]db.Tracker, error) {
	return uc.q.GetTrackersWithPagination(ctx, db.GetTrackersWithPaginationParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
}
