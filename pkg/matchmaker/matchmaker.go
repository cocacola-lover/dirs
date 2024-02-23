package matchmaker

import (
	l "dirs/pkg/logger"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"os"
)

type actualMatchmaker struct {
	// Search requests
	requests map[string][]*dtasks.AskInfoTask
	// Array with stored info
	store map[string]string
}

// Returns processedtrue if info found
// else return false
func (m actualMatchmaker) ProcessAskInfoTask(task *dtasks.AskInfoTask) bool {
	val, ok := m.store[task.Search]

	if ok {
		task.Result = &val
		return true
	} else {
		m.requests[task.Search] = append(m.requests[task.Search], task)
		return false
	}
}

func (m actualMatchmaker) ProcessSortInfoTask(task *dtasks.SortInfoTask) []*dtasks.AskInfoTask {
	m.store[task.Search] = *task.Result

	awaitingProcessing := m.requests[*task.Result]
	m.requests[task.Search] = nil

	for i := range awaitingProcessing {
		awaitingProcessing[i].Result = task.Result
	}

	return awaitingProcessing
}

func NewMatchmaker(Error l.Logger) Matchmaker {

	var knownInfo map[string]string
	marshalErr := json.Unmarshal([]byte(os.Getenv("knownInfo")), &knownInfo)
	if marshalErr != nil {
		Error.Println("Failed to unmarshal KnownInfo")
		return &actualMatchmaker{requests: make(map[string][]*dtasks.AskInfoTask), store: make(map[string]string)}
	}

	return &actualMatchmaker{requests: make(map[string][]*dtasks.AskInfoTask), store: knownInfo}
}
