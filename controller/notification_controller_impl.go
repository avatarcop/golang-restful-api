package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type NotificationControllerImpl struct {
	NotificationService service.NotificationService
}

func NewNotificationController(notificationService service.NotificationService) NotificationController {
	return &NotificationControllerImpl{
		NotificationService: notificationService,
	}
}

func (controller *NotificationControllerImpl) Health(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationResponse := controller.NotificationService.Health(request.Context())
	webResponse := notificationResponse

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) HealthCheckDB(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	healthResponses := controller.NotificationService.HealthCheckDB(request.Context())

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   healthResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationFindByIdRequest := web.NotificationFindByIdRequest{}
	helper.ReadFromRequestQueryParam(request, &notificationFindByIdRequest)

	notificationResponses := controller.NotificationService.FindByUserId(request.Context(), notificationFindByIdRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) Save(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationSaveRequest := web.NotificationSaveRequest{}
	helper.ReadFromRequestBody(request, &notificationSaveRequest)

	notificationResponse := controller.NotificationService.Save(request.Context(), notificationSaveRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
