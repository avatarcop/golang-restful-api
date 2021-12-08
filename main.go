package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	notificationRepository := repository.NewNotificationRepository()
	notificationService := service.NewNotificationService(notificationRepository, db, validate)
	notificationController := controller.NewNotificationController(notificationService)
	router := app.NewRouter(notificationController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	log.Print("Success running server at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err, "error main when listen server")
}
