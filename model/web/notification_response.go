package web

import "time"

type NotificationResponse struct {
	NotificationId int       `json:"notification_id"`
	UserId         int       `json:"user_id"`
	Title          string    `json:"title"`
	Type           string    `json:"type"`
	Message        string    `json:"message"`
	IsRead         bool      `json:"is_read"`
	IconImageUrl   string    `json:"icon_image_url"`
	ImageUrl       string    `json:"image_url"`
	SentStatus     string    `json:"sent_status"`
	RequestRaw     string    `json:"request_raw"`
	ResponseRaw    string    `json:"response_raw"`
	DateIn         time.Time `json:"date_in"`
	UserIn         string    `json:"user_in"`
	DateUp         time.Time `json:"date_up"`
	UserUp         string    `json:"user_up"`
	StatusRecord   string    `json:"status_record"`
}
