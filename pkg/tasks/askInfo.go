package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"
)

// Number of stage 2:
// 0 : Not done
// 1 : Done

var id = 1

type AskInfoTask struct {
	id     string
	from   string
	search string
	stage  int
}

func (t AskInfoTask) GetStage() int {
	return t.stage
}

func (t AskInfoTask) IsDone() bool {
	return t.stage == 1
}

func (t AskInfoTask) String() string {
	return fmt.Sprintf("AskInfoTask #%s : search for %s", t.id, t.search)
}

func NewAskInfoTask(r *drequests.AskInfoRequest) AskInfoTask {
	defer func() { id += 1 }()
	return AskInfoTask{from: r.From, search: r.Search, id: fmt.Sprint(id), stage: 0}
}
