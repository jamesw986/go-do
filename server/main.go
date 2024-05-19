package main

import (
	"log"
	"net/http"
)

type Task struct {
	ID int
	Title string
	Body string
	Done bool
}

var tasks []Task

func main() {
	http.HandleFunc("GET /tasks", getAllTasksHandler)
	http.HandleFunc("GET /tasks/{id}", getTaskHandler)
	http.HandleFunc("POST /tasks", postHandler)
	http.HandleFunc("PATCH /tasks/{id}", patchHandler)
	http.HandleFunc("DELETE /tasks/{id}", deleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}