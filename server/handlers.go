package main

import (
	"net/http"
)

func getAllTasksHandler(w http.ResponseWriter, req *http.Request) {
	sendJSONResponse(w, tasks, http.StatusOK)
}

func getTaskHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")

	taskFound := false
	for _, task := range tasks {
		if idsMatch(task.ID, id) {
			taskFound = true
			sendJSONResponse(w, task, http.StatusFound)
			break
		}
	}

	if !taskFound {
		sendJSONResponse(w, "Task does not exist", http.StatusNotFound)
	}
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	newTask := parseJSON[Task](w, req)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	sendJSONResponse(w, newTask, http.StatusCreated)
}

func patchHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")
	patch := parseJSON[Task](w, req)

	for i, task := range tasks {
		if idsMatch(task.ID, id) {
			updateTask(&tasks[i], &patch)
			break
		}
	}

	sendJSONResponse(w, tasks, http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	id := getId(req, "/tasks/")

	for i, task := range tasks {
		if idsMatch(task.ID, id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	sendJSONResponse(w, tasks, http.StatusOK)
}