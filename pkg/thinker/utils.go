package thinker

import dtasks "dirs/pkg/tasks"

func convertArrayToITask[TaskType dtasks.ITask](tasks []TaskType) []dtasks.ITask {
	interfaceArray := make([]dtasks.ITask, len(tasks))

	for i := range tasks {
		interfaceArray = append(interfaceArray, tasks[i])
	}

	return interfaceArray
}
