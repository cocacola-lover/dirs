package matchmaker

import dtasks "dirs/pkg/tasks/baseTask"

type Matchmaker struct {
	// Waiting line for tasks
	tasks []dtasks.ITask
	// Array with stored info
	store map[string]string
}

func (m Matchmaker) SearchStore(id string) (string, bool) {
	val, ok := m.store[id]
	return val, ok
}

func (m *Matchmaker) QueueUp(task dtasks.ITask) {
	m.tasks = append(m.tasks, task)
}
