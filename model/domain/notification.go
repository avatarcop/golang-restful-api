package domain

import "time"

type Notification struct {
	NotificationId int
	UserId         int
	Title          string
	Type           string
	Message        string
	IsRead         bool
	IconImageUrl   string
	ImageUrl       string
	SentStatus     string
	RequestRaw     string
	ResponseRaw    string
	DateIn         time.Time
	UserIn         string
	DateUp         time.Time
	UserUp         string
	StatusRecord   string
}

type NotificationFindById struct {
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
	ResponseRaw    string
	DateIn         time.Time
	UserIn         string
	DateUp         time.Time
	UserUp         string
	StatusRecord   string
}
