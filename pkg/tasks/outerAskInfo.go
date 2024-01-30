package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"
)

type OuterAskInfoTask struct {
	From   string
	Search string
	Result *string
}

func (t OuterAskInfoTask) GetTaskId() DTaskId {
	return OuterAskInfoId
}

func (t OuterAskInfoTask) String() string {
	return fmt.Sprintf("OuterAskInfoTask from %s : search for %s", t.From, t.Search)
}

func NewOuterAskInfoTaskPointer(r drequests.AskInfoRequest) *OuterAskInfoTask {
	task := OuterAskInfoTask{From: r.From, Search: r.Search}
	return &task
}
