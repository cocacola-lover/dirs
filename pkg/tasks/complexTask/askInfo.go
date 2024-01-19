package dtaskscomplex

import (
	"dirs/pkg/matchmaker"
	drequests "dirs/pkg/requests"
	. "dirs/pkg/tasks/baseTask"
	"fmt"
	"log"
)

// Number of stage 2:
// 0 : Awaits sorting
// 1 : Awaits broadcasting

var id = 1

type AskInfoTask struct {
	id     string
	from   string
	search string
	result *string
	stage  int
}

// Getters

func (t AskInfoTask) Stage() int {
	return t.stage
}

func (t AskInfoTask) Search() string {
	return t.search
}

func (t AskInfoTask) NextStop() DTaskStop {
	switch t.stage {
	case 0:
		return MatchmakerStop
	case 1:
		return BroadcasterStop
	}
	log.Fatalf("%s has moved to unregistered stage", t.String())
	return 0
}

func (t AskInfoTask) String() string {
	return fmt.Sprintf("AskInfoTask #%s : search for %s", t.id, t.search)
}

// Setters

func (t *AskInfoTask) UseMatchmaker(matchmaker matchmaker.Matchmaker) *AskInfoTask {
	if t.NextStop() != MatchmakerStop {
		log.Fatalf("%s has moved to matchmaker at the wrong stage", t.String())
		return nil
	}

	if res, ok := matchmaker.SearchStore(t.id); ok {
		t.result = &res
		t.stage++
		return t
	} else {
		matchmaker.QueueUp(t)
		return nil
	}

}

func NewAskInfoTask(r *drequests.AskInfoRequest) AskInfoTask {
	defer func() { id += 1 }()
	return AskInfoTask{from: r.From, search: r.Search, id: fmt.Sprint(id), stage: 0}
}
