package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(notificationController controller.NotificationController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/health", notificationController.Health)
	router.GET("/health_check_db", notificationController.HealthCheckDB)
	router.GET("/v1/notification", notificationController.FindByUserId)
	router.POST("/v1/notification/save", notificationController.Save)
	router.POST("/v1/notification/update", notificationController.Update)
	router.POST("/v1/notification/update/read", notificationController.UpdateRead)
	router.POST("/v1/notification/delete/soft", notificationController.DeleteSoft)
	router.POST("/v1/notification/delete/hard", notificationController.DeleteHard)
	router.POST("/v1/notification/send", notificationController.Send)

	router.PanicHandler = exception.ErrorHandler

	return router
}
