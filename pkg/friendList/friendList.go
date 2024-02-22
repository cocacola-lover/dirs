package friendlist

import (
	"dirs/pkg/logger"
	"encoding/json"
	"os"
)

type FriendList struct {
	Friends []string
}

func (l FriendList) IsFriend(url string) bool {
	check := false

	for _, v := range l.Friends {
		if v == url {
			check = true
		}
	}

	return check
}

func NewFriendList(logger logger.Logger) FriendList {

	var friendsArr []string
	marshalErr := json.Unmarshal([]byte(os.Getenv("friends")), &friendsArr)
	if marshalErr != nil {
		logger.Error.Println("Failed to unmarshal FriendList")
		return FriendList{Friends: make([]string, 0)}
	}

	return FriendList{Friends: friendsArr}
}
