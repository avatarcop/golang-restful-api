package app

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// func NewDB() *sql.DB {
// 	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api")
// 	helper.PanicIfError(err)
// 	log.Output(0, "No Error DB found")
// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	db.SetConnMaxIdleTime(10 * time.Minute)

// 	return db
// }

func NewDB() *sql.DB {
	// host := os.Getenv("POSTGRES_HOST")
	// user := os.Getenv("POSTGRES_USER")
	// password := os.Getenv("POSTGRES_PASSWORD")
	// databaseName := os.Getenv("POSTGRES_DB_NAME")

	// desc := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, databaseName)
	desc := fmt.Sprintf("host=localhost port=5432 user=postgres password=123 dbname=lokal_belajar sslmode=disable")
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
