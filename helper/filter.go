package helper

type FilterNotificationFindByUserId struct {
	UserId       int
	Page         int
	ItemsPerPage int
	Type         string
	StatusRecord string
	IsRead       string
}
