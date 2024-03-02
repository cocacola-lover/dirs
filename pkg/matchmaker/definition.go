package matchmaker

import t "dirs/pkg/tasks"

type Matchmaker interface {
	// Processes AskInfoTask
	//
	// Matchmaker determines whether info is available
	// and if request has already been registered
	//
	// If info is available, then Matchmaker puts info into task
	// and returns true as the first argument,
	// otherwise Matchmaker puts task in waiting queue
	// and return false as the first argument.
	//
	// If request has already been registered
	// then Matchmaker returns true as the second argument,
	// otherwise returns false as the second argument
	ProcessAskInfoTask(task *t.AskInfoTask) (bool, bool)
	ProcessSortInfoTask(task *t.SortInfoTask) []*t.AskInfoTask
}
