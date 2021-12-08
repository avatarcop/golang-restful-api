package web

type NotificationSendRequest struct {
	UserId  int    `validate:"required" json:"user_id"`
	Title   string `validate:"required" json:"title"`
	Message string `validate:"required" json:"message"`
}
