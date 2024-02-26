package broadcaster

import (
	envp "dirs/pkg/environment"
	rp "dirs/pkg/requests"
	tp "dirs/pkg/tasks"
	"encoding/json"
	"net/http"
	"os"
)

func ProcessAskInfoTask(task *tp.AskInfoTask) error {
	jsonBody, err := json.Marshal(&rp.SendInfoRequest{Search: task.Search, Info: *task.Result})
	if err != nil {
		return err
	}

	return sendRequest(jsonBody, task.From, "/send", http.MethodPost)
}

func ProcessDemandInfoTask(task *tp.DemandInfoTask, env envp.Environment) error {
	jsonBody, err := json.Marshal(&rp.AskInfoRequest{Search: task.Search, From: os.Getenv("address")})
	if err != nil {
		return err
	}

	for _, friend := range env.FriendList.Friends() {
		sendRequest(jsonBody, friend, "/ask", http.MethodPost)
	}

	return nil
}
