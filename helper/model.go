package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/web"
	"strconv"
)

func ToNotificationResponse(notification domain.Notification) web.NotificationResponse {
	isRead, err := strconv.ParseBool(notification.IsRead)
	PanicIfError(err)
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
		ResponsesRaw:   notification.ResponsesRaw,
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

func ToHealthCheckDBResponse(notification domain.Notification) web.HealthCheckDBResponse {
	return web.HealthCheckDBResponse{
		NotificationId: notification.NotificationId,
	}
}
