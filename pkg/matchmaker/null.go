package matchmaker

import t "dirs/pkg/tasks"

type nullMatchmaker struct{}

func (nm nullMatchmaker) ProcessAskInfoTask(task *t.AskInfoTask) bool { return true }
func (nm nullMatchmaker) ProcessSortInfoTask(task *t.SortInfoTask) []*t.AskInfoTask {
	return make([]*t.AskInfoTask, 0)
}

func NullMatchmaker() Matchmaker {
	return nullMatchmaker{}
}
