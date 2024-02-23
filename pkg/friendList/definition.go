package friendlist

type FriendList interface {
	Friends() []string
	IsFriend(url string) bool
}
