package friendlist

import (
	"dirs/pkg/logger"
	"encoding/json"
	"os"
)

type actualFriendList struct {
	friends []string
}

func (l actualFriendList) Friends() []string {
	return l.friends
}

func (l actualFriendList) IsFriend(url string) bool {
	check := false

	for _, v := range l.friends {
		if v == url {
			check = true
		}
	}

	return check
}

func NewFriendList(Info, Error logger.Logger) FriendList {

	var friendsArr []string
	marshalErr := json.Unmarshal([]byte(os.Getenv("friends")), &friendsArr)
	if marshalErr != nil {
		Error.Println("Failed to unmarshal FriendList")
		return &actualFriendList{friends: make([]string, 0)}
	}

	Info.Println("Unmarshal FriendList : ", friendsArr)
	return &actualFriendList{friends: friendsArr}
}
