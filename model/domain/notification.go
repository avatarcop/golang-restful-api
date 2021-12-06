package domain

import "time"

type Notification struct {
	NotificationId int
	UserId         int
	Title          string
	Type           string
	Message        string
	IsRead         string
	IconImageUrl   string
	ImageUrl       string
	SentStatus     string
	RequestRaw     string
	ResponsesRaw   string
	DateIn         time.Time
	UserIn         string
	DateUp         time.Time
	UserUp         string
	StatusRecord   string
}

type NotificationFindByUserId struct {
	UserId       int
	Page         int
	ItemsPerPage int
	Type         string
	StatusRecord string
	IsRead       string
}
