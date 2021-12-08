package web

type NotificationDeleteSoftRequest struct {
	NotificationId int `validate:"required" json:"notification_id"`
}
