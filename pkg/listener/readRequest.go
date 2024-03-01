package listener

import (
	tp "dirs/pkg/tasks"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readRequest[RequestType interface{}](w http.ResponseWriter, r *http.Request, method string) (*RequestType, error) {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil, _WrongMethodError{Message: fmt.Sprintf("Method  %s is not allowed", r.Method)}
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return nil, readErr
	}
	defer r.Body.Close()

	var requestBody RequestType
	marshalErr := json.Unmarshal(body, &requestBody)
	if marshalErr != nil {
		http.Error(w, "Wrong json", http.StatusBadRequest)
		return nil, marshalErr
	}

	return &requestBody, nil
}

func readRequestAndCreateTask[RequestType interface{}, TaskType tp.ITask](w http.ResponseWriter, r *http.Request, method string, initTask func(RequestType) TaskType) (*TaskType, error) {
	req, err := readRequest[RequestType](w, r, method)
	if err != nil {
		return nil, err
	}

	newTask := initTask(*req)
	return &newTask, nil
}
