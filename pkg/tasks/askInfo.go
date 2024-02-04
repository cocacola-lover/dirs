package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"
)

type AskInfoTask struct {
	From   string
	Search string
	Result *string
}

func (t AskInfoTask) GetTaskId() DTaskId {
	return AskInfoId
}

func (t AskInfoTask) String() string {
	return fmt.Sprintf("AskInfoTask from %s : search for %s", t.From, t.Search)
}

func NewAskInfoTaskPointer(r drequests.AskInfoRequest) *AskInfoTask {
	task := AskInfoTask{From: r.From, Search: r.Search}
	return &task
}
