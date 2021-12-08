package web

type NotificationDeleteHardRequest struct {
	NotificationId int `validate:"required" json:"notification_id"`
}
