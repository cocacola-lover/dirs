package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"
)

type SortInfoTask struct {
	Search string
	Result *string
}

func (t SortInfoTask) GetTaskId() DTaskId {
	return SortInfoId
}

func (t SortInfoTask) String() string {
	return fmt.Sprintf("SortInfoTask for %s : %s", t.Search, t.Search)
}

func NewSortInfoTaskPointer(r drequests.SendInfoRequest) *SortInfoTask {
	task := SortInfoTask{Result: &r.Info, Search: r.Search}
	return &task
}
