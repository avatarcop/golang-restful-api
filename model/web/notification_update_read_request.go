package web

type NotificationUpdateReadRequest struct {
	NotificationId int  `validate:"required" json:"notification_id"`
	IsRead         bool `json:"is_read"`
}
