package email

import (
	"context"

	db "github.com/corentings/email-tracker/db/sqlc"
	"github.com/google/uuid"
)

type IUseCase interface {
	GetEmail(ctx context.Context, imageUUID uuid.UUID) (db.Email, error)
	AddTracking(ctx context.Context, email db.Email, ip string) error
	CreateEmail(ctx context.Context, recipient, subject string) (db.Email, error)
	GetEmailsWithPagination(ctx context.Context, limit int, offset int) ([]db.Email, error)
	GetTrackersWithPagination(ctx context.Context, limit int, offset int) ([]db.Tracker, error)
}
