package broadcaster

import (
	drequests "dirs/pkg/requests"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"net/http"
	"os"
)

func ProcessAskInfoTask(task *dtasks.AskInfoTask, serviceStore ss.ServiceStore) error {
	jsonBody, err := json.Marshal(&drequests.SendInfoRequest{Search: task.Search, Info: *task.Result})
	if err != nil {
		serviceStore.Logger.Error.Printf("client: could not marshal request: %s\n", err)
		return err
	}

	return sendRequest(jsonBody, task.From, http.MethodPost, serviceStore)
}

func ProcessDemandInfoTask(task *dtasks.DemandInfoTask, serviceStore ss.ServiceStore) error {
	jsonBody, err := json.Marshal(&drequests.AskInfoRequest{Search: task.Search, From: os.Getenv("address")})
	if err != nil {
		serviceStore.Logger.Error.Printf("client: could not marshal request: %s\n", err)
		return err
	}

	for _, friend := range serviceStore.FriendList.Friends {
		sendRequest(jsonBody, friend, http.MethodPost, serviceStore)
	}

	return nil
}
