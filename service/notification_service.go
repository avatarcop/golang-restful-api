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
	Update(ctx context.Context, request web.NotificationUpdateRequest) web.NotificationResponse
	UpdateRead(ctx context.Context, request web.NotificationUpdateReadRequest) web.NotificationResponse
	DeleteSoft(ctx context.Context, request web.NotificationDeleteSoftRequest) web.NotificationResponse
	DeleteHard(ctx context.Context, request web.NotificationDeleteHardRequest)
	Send(ctx context.Context, request web.NotificationSendRequest) web.NotificationResponse
}
