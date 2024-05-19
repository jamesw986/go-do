package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

func sendJSONResponse[T any](w http.ResponseWriter, body T, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "Failed to encode response body", http.StatusInternalServerError)
	}
}

func parseJSON[T any](w http.ResponseWriter, req *http.Request) T {
	var parsedBody T
	if err := json.NewDecoder(req.Body).Decode(&parsedBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	return parsedBody
}

func updateTask(originalTask *Task, patch *Task) {
	originalTaskValues := reflect.ValueOf(originalTask).Elem()
	patchValues := reflect.ValueOf(patch).Elem()

	for i := 0; i < originalTaskValues.NumField(); i++ {
		originalTaskField := originalTaskValues.Field(i)
		patchField := patchValues.Field(i)

		if !patchField.IsZero() {
			originalTaskField.Set(patchField)
		}
	}
}

func getId(req *http.Request, pathToId string) string {
	return req.URL.Path[len(pathToId):]
}

func idsMatch(taskId int, requestId string) bool {
	parsedId, _ := strconv.Atoi(requestId)
	return taskId == parsedId
}