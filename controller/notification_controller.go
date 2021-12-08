package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NotificationController interface {
	Health(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	HealthCheckDB(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateRead(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteSoft(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteHard(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Send(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
