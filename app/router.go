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
	// masih error
	router.GET("/v1/notification", notificationController.FindByUserId)
	// ------
	router.POST("/v1/notification/save", notificationController.Save)

	router.PanicHandler = exception.ErrorHandler

	return router
}
