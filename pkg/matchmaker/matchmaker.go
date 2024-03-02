package matchmaker

import (
	l "dirs/pkg/logger"
	dt "dirs/pkg/tasks"
	"encoding/json"
	"os"
)

type actualMatchmaker struct {
	// Search requests by id
	requests map[string]*dt.AskInfoTask
	// Array with stored info
	store map[string]string
}

func (m actualMatchmaker) ProcessAskInfoTask(task *dt.AskInfoTask) (bool, bool) {
	val, ok := m.store[task.Search]

	if ok {
		task.Result = &val
		return true, false
	} else {
		// Check if request is already registered
		_, isRegistered := m.requests[task.Id]
		// Return if already registered
		if isRegistered {
			return false, true
		}

		// Add to queue
		m.requests[task.Id] = task
		return false, false
	}
}

func (m actualMatchmaker) ProcessSortInfoTask(task *dt.SortInfoTask) []*dt.AskInfoTask {
	m.store[task.Search] = *task.Result

	awaitingProcessing := []*dt.AskInfoTask{}

	for key, vTask := range m.requests {
		if vTask.Search == task.Search {
			// Answer AskInfo
			vTask.Result = task.Result

			awaitingProcessing = append(awaitingProcessing, vTask)
			// Remove from queue
			delete(m.requests, key)
		}
	}

	return awaitingProcessing
}

func NewMatchmaker(Error l.Logger) Matchmaker {

	var knownInfo map[string]string
	marshalErr := json.Unmarshal([]byte(os.Getenv("knownInfo")), &knownInfo)
	if marshalErr != nil {
		Error.Println("Failed to unmarshal KnownInfo")
		return &actualMatchmaker{requests: make(map[string]*dt.AskInfoTask), store: make(map[string]string)}
	}

	return &actualMatchmaker{requests: make(map[string]*dt.AskInfoTask), store: knownInfo}
}
