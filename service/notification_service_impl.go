package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/model/domain"
	"golang-restful-api/model/entity"
	"golang-restful-api/model/web"
	"golang-restful-api/repository"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	helper.PanicIfError(err, "error notification service at func healthcheckdb when begin DB")
	defer helper.CommitOrRollback(tx)

	healthCheckDBs := service.NotificationRepository.HealthCheckDB(ctx, tx)

	return helper.ToHealthCheckDBResponse(healthCheckDBs)
}

func (service *NotificationServiceImpl) FindByUserId(ctx context.Context, request web.NotificationFindByIdRequest) []web.NotificationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "error notification service at func findbyuserid when validate request")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func findbyuserid when DB begin")

	defer helper.CommitOrRollback(tx)

	filterFindByUserId := helper.FilterNotificationFindByUserId{
		UserId:       request.UserId,
		Page:         request.Page,
		ItemsPerPage: request.ItemsPerPage,
		Type:         request.Type,
		StatusRecord: request.StatusRecord,
		IsRead:       request.IsRead,
	}

	notification := service.NotificationRepository.FindByUserId(ctx, tx, filterFindByUserId)

	return helper.ToNotificationFindByIdResponses(notification)
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

func (service *NotificationServiceImpl) Update(ctx context.Context, request web.NotificationUpdateRequest) web.NotificationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "error notification service at func update when validate request")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func update when DB begin")

	defer helper.CommitOrRollback(tx)

	notification, err := service.NotificationRepository.FindOneByNotificationId(ctx, tx, request.NotificationId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	notification.UserId = request.UserId
	notification.Type = strings.ToTitle(request.Type)
	notification.Title = request.Title
	notification.Message = request.Message
	notification.IsRead = request.IsRead
	notification.IconImageUrl = request.IconImageUrl
	notification.ImageUrl = request.ImageUrl
	notification.StatusRecord = "A"
	notification.DateUp = time.Now()
	notification.UserUp = "-"

	notification = service.NotificationRepository.Update(ctx, tx, notification)

	return helper.ToNotificationResponse(notification)
}

func (service *NotificationServiceImpl) UpdateRead(ctx context.Context, request web.NotificationUpdateReadRequest) web.NotificationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "error notification service at func updateread when validate request")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func updateread when DB begin")

	defer helper.CommitOrRollback(tx)

	notification, err := service.NotificationRepository.FindOneByNotificationId(ctx, tx, request.NotificationId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	notification.IsRead = request.IsRead
	notification.StatusRecord = "U"
	notification.DateUp = time.Now()
	notification.UserUp = "-"

	notification = service.NotificationRepository.UpdateRead(ctx, tx, notification)

	return helper.ToNotificationResponse(notification)
}

func (service *NotificationServiceImpl) DeleteSoft(ctx context.Context, request web.NotificationDeleteSoftRequest) web.NotificationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "error notification service at func deletesoft when validate request")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func deletesoft when DB begin")
	defer helper.CommitOrRollback(tx)

	notification, err := service.NotificationRepository.FindOneByNotificationId(ctx, tx, request.NotificationId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	notification.StatusRecord = "D"
	notification.DateUp = time.Now()
	notification.UserUp = "-"

	notification = service.NotificationRepository.DeleteSoft(ctx, tx, notification)

	return helper.ToNotificationResponse(notification)
}

func (service *NotificationServiceImpl) DeleteHard(ctx context.Context, request web.NotificationDeleteHardRequest) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func delete hard when DB begin")

	defer helper.CommitOrRollback(tx)

	notification, err := service.NotificationRepository.FindOneByNotificationId(ctx, tx, request.NotificationId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.NotificationRepository.DeleteHard(ctx, tx, notification)
}

func (service *NotificationServiceImpl) Send(ctx context.Context, request web.NotificationSendRequest) web.NotificationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err, "error notification service at func send when validate request")

	tx, err := service.DB.Begin()
	helper.PanicIfError(err, "error notification service at func send when DB begin")

	defer helper.CommitOrRollback(tx)

	notification := domain.Notification{
		UserId:       request.UserId,
		Title:        request.Title,
		Message:      request.Message,
		DateIn:       time.Now(),
		UserIn:       "-",
		StatusRecord: "A",
		SentStatus:   "process",
	}

	// save log notif
	logNotification := service.NotificationRepository.Save(ctx, tx, notification)
	helper.PanicIfError(err, "error notification service at func send when save log notif")

	notifThirdParty := os.Getenv("THIRD_PARTY_NOTIF_ACTIVE")

	if notifThirdParty == "onesignal" {
		url := os.Getenv("THIRD_PARTY_NOTIF_ONE_SIGNAL_APP_URL")
		method := "POST"

		payload := entity.OnesignalPayload{
			AppId:                     os.Getenv("THIRD_PARTY_NOTIF_ONE_SIGNAL_APP_ID"),
			IncludeExternalUserIds:    []string{strconv.Itoa(request.UserId)},
			ChannelForExternalUserIds: "push",
			Headings: entity.Headings{
				En: request.Title,
			},
			Contents: entity.Contents{
				En: request.Message,
			},
		}

		payloadEncoded, _ := json.Marshal(&payload)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, bytes.NewReader(payloadEncoded))

		status_sent := "success"

		if err != nil {
			log.Fatal("error notification service at func update when send notif to one signal")
			log.Fatal(err)
			status_sent = "failed"
		}
		req.Header.Add("Content-Type", "application/json; charset=utf-8")
		req.Header.Add("Authorization", "Basic "+os.Getenv("THIRD_PARTY_NOTIF_ONE_SIGNAL_APP_KEY"))

		res, err := client.Do(req)
		if err != nil {
			log.Fatal("error notification service at func update when read response")
			log.Fatal(err)
			status_sent = "failed"
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal("error notification service at func update when get body response")
			log.Fatal(err)
			status_sent = "failed"
		}

		notification = domain.Notification{
			NotificationId: logNotification.NotificationId,
			UserId:         logNotification.UserId,
			Title:          logNotification.Title,
			Type:           logNotification.Type,
			Message:        logNotification.Message,
			IsRead:         logNotification.IsRead,
			IconImageUrl:   logNotification.IconImageUrl,
			ImageUrl:       logNotification.ImageUrl,
			DateIn:         logNotification.DateIn,
			UserIn:         logNotification.UserIn,
			SentStatus:     status_sent,
			RequestRaw:     string(payloadEncoded[:]),
			ResponseRaw:    string(body[:]),
			StatusRecord:   "U",
			DateUp:         time.Now(),
			UserUp:         "-",
		}

		// update log notif
		notification = service.NotificationRepository.Update(ctx, tx, notification)
	}
	return helper.ToNotificationResponse(notification)

}
