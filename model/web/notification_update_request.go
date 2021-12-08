package web

type NotificationUpdateRequest struct {
	NotificationId int    `validate:"required" json:"notification_id"`
	UserId         int    `validate:"required" json:"user_id"`
	Type           string `validate:"required" json:"type"`
	Title          string `validate:"required" json:"title"`
	Message        string `validate:"required" json:"message"`
	IsRead         bool   `json:"is_read"`
	IconImageUrl   string `json:"icon_image_url"`
	ImageUrl       string `json:"image_url"`
}
