package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"strconv"
)

func ToNotificationResponse(notification domain.Notification) web.NotificationResponse {
	return web.NotificationResponse{
		NotificationId: notification.NotificationId,
		UserId:         notification.UserId,
		Title:          notification.Title,
		Type:           notification.Type,
		Message:        notification.Message,
		IsRead:         notification.IsRead,
		IconImageUrl:   notification.IconImageUrl,
		ImageUrl:       notification.ImageUrl,
		SentStatus:     notification.SentStatus,
		RequestRaw:     notification.RequestRaw,
		ResponseRaw:    notification.ResponseRaw,
		DateIn:         notification.DateIn,
		UserIn:         notification.UserIn,
		DateUp:         notification.DateUp,
		UserUp:         notification.UserUp,
		StatusRecord:   notification.StatusRecord,
	}
}

func ToNotificationResponses(notifications []domain.Notification) []web.NotificationResponse {
	var notificationResponses []web.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, ToNotificationResponse(notification))
	}
	return notificationResponses
}

func ToNotificationFindByIdResponse(notification domain.NotificationFindById) web.NotificationResponse {
	isRead, err := strconv.ParseBool(notification.IsRead)
	PanicIfError(err, "error helper/model at func ToNotificationFindByIdResponse when convert is read to bool")
	return web.NotificationResponse{
		NotificationId: notification.NotificationId,
		UserId:         notification.UserId,
		Title:          notification.Title,
		Type:           notification.Type,
		Message:        notification.Message,
		IsRead:         isRead,
		IconImageUrl:   notification.IconImageUrl,
		ImageUrl:       notification.ImageUrl,
		SentStatus:     notification.SentStatus,
		RequestRaw:     notification.RequestRaw,
		ResponseRaw:    notification.ResponseRaw,
		DateIn:         notification.DateIn,
		UserIn:         notification.UserIn,
		DateUp:         notification.DateUp,
		UserUp:         notification.UserUp,
		StatusRecord:   notification.StatusRecord,
	}
}

func ToNotificationFindByIdResponses(notifications []domain.NotificationFindById) []web.NotificationResponse {
	var notificationResponses []web.NotificationResponse
	for _, notification := range notifications {
		notificationResponses = append(notificationResponses, ToNotificationFindByIdResponse(notification))
	}
	return notificationResponses
}

func ToHealthCheckDBResponse(notification domain.Notification) web.HealthCheckDBResponse {
	return web.HealthCheckDBResponse{
		NotificationId: notification.NotificationId,
	}
}
