package matchmaker

import dtasks "dirs/pkg/tasks"

type Matchmaker struct {
	// Search requests
	requests map[string][]*dtasks.OuterAskInfoTask
	// Array with stored info
	store map[string]string
}

// Returns processedtrue if info found
// else return false
func (m Matchmaker) ProcessOuterAskInfoTask(task *dtasks.OuterAskInfoTask) bool {
	val, ok := m.store[task.Search]

	if ok {
		task.Result = &val
		return true
	} else {
		m.requests[task.Search] = append(m.requests[task.Search], task)
		return false
	}
}

func (m Matchmaker) ProcessSortInfoTask(task *dtasks.SortInfoTask) []*dtasks.OuterAskInfoTask {
	m.store[task.Search] = *task.Result

	awaitingProcessing := m.requests[*task.Result]
	m.requests[task.Search] = nil

	for i := range awaitingProcessing {
		awaitingProcessing[i].Result = task.Result
	}

	return awaitingProcessing
}

func NewMatchmaker() Matchmaker {
	return Matchmaker{requests: make(map[string][]*dtasks.OuterAskInfoTask), store: make(map[string]string)}
}
