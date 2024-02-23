package matchmaker

import t "dirs/pkg/tasks"

type Matchmaker interface {
	ProcessAskInfoTask(task *t.AskInfoTask) bool
	ProcessSortInfoTask(task *t.SortInfoTask) []*t.AskInfoTask
}
