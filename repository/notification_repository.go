package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/model/domain"
)

type NotificationRepository interface {
	Health(ctx context.Context, tx *sql.Tx) int
	HealthCheckDB(ctx context.Context, tx *sql.Tx) domain.Notification
	FindByUserId(ctx context.Context, tx *sql.Tx, notification domain.NotificationFindByUserId) []domain.Notification
	Save(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification
}
