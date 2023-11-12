package economy

type Economy interface {
	GetPackage(id string) (string bool)
	GetWants() map[string]bool
	GetFriends() map[string]bool

	Need(id string)
	Put(id, value string)
}
