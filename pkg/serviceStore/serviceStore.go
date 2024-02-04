package servicestore

import (
	friendlist "dirs/pkg/friendList"
	"dirs/pkg/matchmaker"
	dtasks "dirs/pkg/tasks"
)

// Struct holding all instances used by components
type ServiceStore struct {
	Matchmaker *matchmaker.Matchmaker
	FriendList *friendlist.FriendList
	TaskCh     *chan dtasks.ITask
}
