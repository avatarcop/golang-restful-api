package service

import (
	"context"
	"golang-restful-api/model/web"
)

type NotificationService interface {
	Health(ctx context.Context) web.HealthResponse
	HealthCheckDB(ctx context.Context) web.HealthCheckDBResponse
	FindByUserId(ctx context.Context, request web.NotificationFindByIdRequest) []web.NotificationResponse
	Save(ctx context.Context, request web.NotificationSaveRequest) web.NotificationResponse
}
