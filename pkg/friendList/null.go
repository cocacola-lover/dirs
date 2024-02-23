package friendlist

type nullFriendList struct{}

func (l nullFriendList) Friends() []string {
	return make([]string, 0)
}
func (l nullFriendList) IsFriend(url string) bool {
	return false
}

func NullFriendList() FriendList {
	return nullFriendList{}
}
