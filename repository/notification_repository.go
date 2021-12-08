package repository

import (
	"context"
	"database/sql"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
)

type NotificationRepository interface {
	Health(ctx context.Context, tx *sql.Tx) int
	HealthCheckDB(ctx context.Context, tx *sql.Tx) domain.Notification
	FindByUserId(ctx context.Context, tx *sql.Tx, notification helper.FilterNotificationFindByUserId) []domain.NotificationFindById
	Save(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification
	FindOneByNotificationId(ctx context.Context, tx *sql.Tx, notificationId int) (domain.Notification, error)
	Update(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification
	UpdateRead(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification
	DeleteSoft(ctx context.Context, tx *sql.Tx, notification domain.Notification) domain.Notification
	DeleteHard(ctx context.Context, tx *sql.Tx, notification domain.Notification)
}
