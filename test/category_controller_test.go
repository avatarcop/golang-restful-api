package test

import (
	"database/sql"
	"fmt"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func testSetupTestDB() *sql.DB {
	// host := os.Getenv("POSTGRES_HOST")
	// user := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// databaseName := os.Getenv("POSTGRES_DB_NAME")

	// desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)
	desc := fmt.Sprintf("host=172.18.0.2 port=5432 user=postgres password=123 dbname=lokal_renosdb_test sslmode=disable")
	db, err := sql.Open("postgres", desc)

	if err != nil {
		log.Print("Cannot connect to postgree database")
		log.Fatal("This is the error: ", err)
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	log.Print("Success open connection DB")
	return db
}

func createConnectionTest(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db, nil
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	notificationRepository := repository.NewNotificationRepository()
	notificationService := service.NewNotificationService(notificationRepository, db, validate)
	notificationController := controller.NewNotificationController(notificationService)
	router := app.NewRouter(notificationController)

	return middleware.NewAuthMiddleware(router)
}

func truncateNotification(db *sql.DB) {
	db.Exec("TRUNCATE rns_notification")
}

func TestHealthNotificationSuccess(t *testing.T) {
	db := testSetupTestDB()
	truncateNotification(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/health", nil)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
