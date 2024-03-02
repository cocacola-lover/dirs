package broadcaster

import (
	fl "dirs/pkg/friendList"
	rp "dirs/pkg/requests"
	"encoding/json"
	"net/http"
)

func askInfoFromFriends(task rp.AskInfoRequest, friends fl.FriendList) error {
	jsonBody, err := json.Marshal(&task)
	if err != nil {
		return err
	}

	for _, friend := range friends.Friends() {
		sendRequest(jsonBody, friend, "/ask", http.MethodPost)
	}

	return nil
}
