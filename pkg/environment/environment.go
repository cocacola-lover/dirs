package environment

import (
	fl "dirs/pkg/friendList"
	l "dirs/pkg/logger"
	m "dirs/pkg/matchmaker"
)

type Environment struct {
	Info       l.Logger
	Warning    l.Logger
	Error      l.Logger
	Matchmaker m.Matchmaker
	FriendList fl.FriendList
}

func NewEnvironment() Environment {
	i, w, e := l.NewLogger()
	matchmaker := m.NewMatchmaker(e)
	friendList := fl.NewFriendList(e)
	return Environment{Info: i, Warning: w, Error: e, Matchmaker: matchmaker, FriendList: friendList}
}

func NullEnvironment() Environment {
	i, w, e := l.NullLogger()
	matchmaker := m.NullMatchmaker()
	friendlist := fl.NullFriendList()
	return Environment{Info: i, Warning: w, Error: e, Matchmaker: matchmaker, FriendList: friendlist}
}
