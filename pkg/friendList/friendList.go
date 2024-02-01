package friendlist

import (
	"encoding/json"
	"fmt"
	"os"
)

type FriendList struct {
	friends []string
}

func (l FriendList) IsFriend(url string) bool {
	check := false

	for _, v := range l.friends {
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
		return FriendList{friends: make([]string, 0)}
	}

	return FriendList{friends: friendsArr}
}
