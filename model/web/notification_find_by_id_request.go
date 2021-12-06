package web

type NotificationFindByIdRequest struct {
	UserId       int    `validate:"required" schema:"user_id"`
	Page         int    `validate:"required" schema:"page"`
	ItemsPerPage int    `validate:"required" schema:"items_per_page"`
	Type         string `schema:"type"`
	StatusRecord string `schema:"status_record"`
	IsRead       string `schema:"is_read"`
}
