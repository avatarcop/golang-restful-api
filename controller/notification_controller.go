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
}
