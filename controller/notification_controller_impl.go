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

func (controller *NotificationControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationUpdateRequest := web.NotificationUpdateRequest{}
	helper.ReadFromRequestBody(request, &notificationUpdateRequest)

	notificationResponse := controller.NotificationService.Update(request.Context(), notificationUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) UpdateRead(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationUpdateReadRequest := web.NotificationUpdateReadRequest{}
	helper.ReadFromRequestBody(request, &notificationUpdateReadRequest)

	notificationResponse := controller.NotificationService.UpdateRead(request.Context(), notificationUpdateReadRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) DeleteSoft(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationDeleteSoftRequest := web.NotificationDeleteSoftRequest{}
	helper.ReadFromRequestBody(request, &notificationDeleteSoftRequest)

	notificationResponse := controller.NotificationService.DeleteSoft(request.Context(), notificationDeleteSoftRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) DeleteHard(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationDeleteHardRequest := web.NotificationDeleteHardRequest{}
	helper.ReadFromRequestBody(request, &notificationDeleteHardRequest)

	controller.NotificationService.DeleteHard(request.Context(), notificationDeleteHardRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   make([]interface{}, 0),
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *NotificationControllerImpl) Send(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	notificationSendRequest := web.NotificationSendRequest{}
	helper.ReadFromRequestBody(request, &notificationSendRequest)

	notificationResponse := controller.NotificationService.Send(request.Context(), notificationSendRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   notificationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
