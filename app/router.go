package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// router.GET("/v1/buyer/transaction/cart", cartController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
