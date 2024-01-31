package thinker

import (
	"dirs/pkg/broadcaster"
	"dirs/pkg/listener"
	m "dirs/pkg/matchmaker"
	dtasks "dirs/pkg/tasks"
	"fmt"
)

func InitThinker() {
	matchmaker := m.NewMatchmaker()

	taskCh := make(chan dtasks.ITask)

	go listener.Listen(taskCh)

	resolveTasks(taskCh, &matchmaker)
}

func resolveTasks(taskCh chan dtasks.ITask, matchmaker *m.Matchmaker) {
	for {
		task, ok := <-taskCh

		var newTasks []dtasks.ITask

		if !ok {
			fmt.Println("Channel closed")
			return
		}

		switch task.GetTaskId() {
		case dtasks.OuterAskInfoId:
			newTasks = resolveOuterAskInfo(task.(*dtasks.OuterAskInfoTask), matchmaker)
		case dtasks.SortInfoId:
			newTasks = resolveSortInfo(task.(*dtasks.SortInfoTask), matchmaker)
		default:
			fmt.Println("Uknown task")
		}

		for _, v := range newTasks {
			taskCh <- v
		}
	}
}

func resolveOuterAskInfo(task *dtasks.OuterAskInfoTask, matchmaker *m.Matchmaker) []dtasks.ITask {

	if task.Result == nil {
		if !matchmaker.ProcessOuterAskInfoTask(task) {
			return nil
		}
	}

	broadcaster.ProcessOuterAskInfoTask(task)
	return nil
}

func resolveSortInfo(task *dtasks.SortInfoTask, matchmaker *m.Matchmaker) []dtasks.ITask {
	return convertArrayToITask(matchmaker.ProcessSortInfoTask(task))
}

func convertArrayToITask[TaskType dtasks.ITask](tasks []TaskType) []dtasks.ITask {
	interfaceArray := make([]dtasks.ITask, len(tasks))

	for i := range tasks {
		interfaceArray = append(interfaceArray, tasks[i])
	}

	return interfaceArray
}
