package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"
)

type DemandInfoTask struct {
	Search string
}

func (t DemandInfoTask) GetTaskId() DTaskId {
	return DemandInfoId
}

func (t DemandInfoTask) String() string {
	return fmt.Sprintf("DemandInfoTask : %s", t.Search)
}

func NewDemandInfoTaskPointer(r drequests.DemandInfoRequest) *DemandInfoTask {
	task := DemandInfoTask{Search: r.Search}
	return &task
}
