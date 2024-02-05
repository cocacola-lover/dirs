package broadcaster

import (
	drequests "dirs/pkg/requests"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func ProcessAskInfoTask(task *dtasks.AskInfoTask) error {
	jsonBody, err := json.Marshal(&drequests.SendInfoRequest{Search: task.Search, Info: *task.Result})
	if err != nil {
		fmt.Printf("client: could not marshal request: %s\n", err)
		return err
	}

	return sendRequest(jsonBody, task.From, http.MethodPost)
}

func ProcessDemandInfoTask(task *dtasks.DemandInfoTask, serviceStore ss.ServiceStore) error {
	jsonBody, err := json.Marshal(&drequests.AskInfoRequest{Search: task.Search, From: os.Getenv("address")})
	if err != nil {
		fmt.Printf("client: could not marshal request: %s\n", err)
		return err
	}

	for _, friend := range serviceStore.FriendList.Friends {
		sendRequest(jsonBody, friend, http.MethodPost)
	}

	return nil
}
