package listener

import (
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func readRequestAndCreateTask[RequestType interface{}, TaskType dtasks.ITask](w http.ResponseWriter, r *http.Request, method string, initTask func(RequestType) TaskType, key any) error {
	if r.Method != method {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return WrongMethodError{Message: fmt.Sprintf("Method  %s is not allowed", r.Method)}
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return readErr
	}
	defer r.Body.Close()

	var requestBody RequestType
	marshalErr := json.Unmarshal(body, &requestBody)
	if marshalErr != nil {
		http.Error(w, "Wrong json", http.StatusBadRequest)
		return marshalErr
	}

	ctx := r.Context()
	serviceStore := ctx.Value(key).(ss.ServiceStore)
	taskCh := *serviceStore.TaskCh
	taskCh <- initTask(requestBody)

	serviceStore.Logger.Info.Printf("got request %s\n", string(body))

	return nil
}
