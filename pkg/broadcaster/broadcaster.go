package broadcaster

import (
	envp "dirs/pkg/environment"
	rp "dirs/pkg/requests"
	tp "dirs/pkg/tasks"
	"encoding/json"
	"net/http"
	"os"
)

// Processes AskInfoTask
//
// If task includes result -> sends info to the requester.
// Otherwise asks for info from friends.
func ProcessAskInfoTask(task *tp.AskInfoTask, env envp.Environment) error {

	if task.Result == nil {
		return askInfoFromFriends(rp.AskInfoRequest{Search: task.Search, From: os.Getenv("address")}, env.FriendList)
	}

	jsonBody, err := json.Marshal(&rp.SendInfoRequest{Search: task.Search, Info: *task.Result})
	if err != nil {
		return err
	}

	return sendRequest(jsonBody, task.From, "/send", http.MethodPost)
}

func ProcessDemandInfoTask(task *tp.DemandInfoTask, env envp.Environment) error {
	return askInfoFromFriends(rp.AskInfoRequest{Search: task.Search, From: os.Getenv("address")}, env.FriendList)
}
