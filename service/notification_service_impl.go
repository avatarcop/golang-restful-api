package service

import (
	"context"
	"database/sql"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewNotificationService(notificationRepository repository.NotificationRepository, DB *sql.DB, validate *validator.Validate) NotificationService {
	return &NotificationServiceImpl{
		NotificationRepository: notificationRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *NotificationServiceImpl) Health(ctx context.Context) web.HealthResponse {
	return 1
}

func (service *NotificationServiceImpl) HealthCheckDB(ctx context.Context) web.HealthCheckDBResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	healthCheckDBs := service.NotificationRepository.HealthCheckDB(ctx, tx)

	return helper.ToHealthCheckDBResponse(healthCheckDBs)
}

func (service *NotificationServiceImpl) FindByUserId(ctx context.Context, request web.NotificationFindByIdRequest) []web.NotificationResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		log.Fatal("Error validate struct Service Notif FindByUserId:", err)
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Fatal("Error begin DB Service Notif FindByUserId:", err)
		panic(err)
	}
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	filterFindByUserId := domain.NotificationFindByUserId{
		UserId:       request.UserId,
		Page:         request.Page,
		ItemsPerPage: request.ItemsPerPage,
		Type:         request.Type,
		StatusRecord: request.StatusRecord,
		IsRead:       request.IsRead,
	}

	notification := service.NotificationRepository.FindByUserId(ctx, tx, filterFindByUserId)

	return helper.ToNotificationResponses(notification)
}

func (service *NotificationServiceImpl) Save(ctx context.Context, request web.NotificationSaveRequest) web.NotificationResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		log.Fatal("Error validate struct Service Notif Save:", err)
		panic(err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		log.Fatal("Error begin DB Service Notif Save:", err)
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	notification := domain.Notification{
		UserId:       request.UserId,
		Title:        request.Title,
		Type:         request.Type,
		Message:      request.Message,
		IsRead:       request.IsRead,
		IconImageUrl: request.IconImageUrl,
		ImageUrl:     request.ImageUrl,
		SentStatus:   "process",
		DateIn:       time.Now(),
		UserIn:       "manual",
		StatusRecord: "A",
	}

	notification = service.NotificationRepository.Save(ctx, tx, notification)

	return helper.ToNotificationResponse(notification)
}
