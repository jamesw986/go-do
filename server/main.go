package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Task struct {
	ID int
	Title string
	Body string
	Done bool
}

var db *sql.DB


func main() {
	dbConfig := mysql.Config{
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net: "tcp",
		Addr: "127.0.0.1:3307",
		DBName: "godo",
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to DB")

	http.HandleFunc("GET /tasks", getAllTasksHandler)
	http.HandleFunc("GET /tasks/{id}", getTaskHandler)
	http.HandleFunc("POST /tasks", postHandler)
	http.HandleFunc("PUT /tasks/{id}", putHandler)
	http.HandleFunc("DELETE /tasks/{id}", deleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}