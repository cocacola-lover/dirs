package friendlist

import (
	"encoding/json"
	"fmt"
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

func NewFriendList() FriendList {

	var friendsArr []string
	marshalErr := json.Unmarshal([]byte(os.Getenv("friends")), &friendsArr)
	if marshalErr != nil {
		fmt.Print("Failed to unmarshal FriendList")
		return FriendList{Friends: make([]string, 0)}
	}

	return FriendList{Friends: friendsArr}
}
