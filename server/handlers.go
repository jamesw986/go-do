package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func getAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	var tasks []Task

	rows, err := db.Query("SELECT * FROM tasks");
	if err != nil {
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Body, &task.Done); err != nil {
			http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		}
		tasks = append(tasks, task)
	}

	sendJSONResponse(w, tasks, http.StatusOK)
}

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")

	var task Task
	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&task.ID, &task.Title, &task.Body, &task.Done); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task does not exist", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	sendJSONResponse(w, task, http.StatusOK)
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	newTask := parseJSON[Task](w, req)

	result, err := db.Exec("INSERT INTO tasks (title, body, done) VALUES (?, ?, ?)", newTask.Title, newTask.Body, newTask.Done)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusBadRequest)
		return
	}

	newTaskId, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("Failed to retrieve ID of newly created task: %v", err)
	}
	newTask.ID = int(newTaskId)

	sendJSONResponse(w, newTask, http.StatusCreated)
}

func putHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")
	update := parseJSON[Task](w, req)

	_, err := db.Exec("UPDATE tasks SET id = ?, title = ?, body = ?, done = ? WHERE id = ?", id, update.Title, update.Body, update.Done, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Failed to update as task does not exist", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
	}

	sendJSONResponse(w, update, http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")

	if _, err := db.Exec("DELETE FROM tasks WHERE id = ?", id); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Task does not exist", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	responseMessage := fmt.Sprintf("Successfully deleted task with ID %s", id)
	sendJSONResponse(w, responseMessage, http.StatusOK)
}