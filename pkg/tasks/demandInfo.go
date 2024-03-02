package dtasks

import (
	drequests "dirs/pkg/requests"
	"fmt"

	"github.com/google/uuid"
)

type DemandInfoTask struct {
	Id     string
	Search string
}

func (t DemandInfoTask) GetTaskId() DTaskId {
	return DemandInfoId
}

func (t DemandInfoTask) String() string {
	return fmt.Sprintf("DemandInfoTask : %s", t.Search)
}

func NewDemandInfoTaskPointer(r drequests.DemandInfoRequest) *DemandInfoTask {
	task := DemandInfoTask{Search: r.Search, Id: uuid.NewString()}
	return &task
}
