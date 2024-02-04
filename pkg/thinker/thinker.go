package thinker

import (
	"dirs/pkg/broadcaster"
	fl "dirs/pkg/friendList"
	"dirs/pkg/listener"
	m "dirs/pkg/matchmaker"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"fmt"
)

func InitThinker() {
	matchmaker := m.NewMatchmaker()
	taskCh := make(chan dtasks.ITask)
	friendList := fl.NewFriendList()

	serviceStore := ss.ServiceStore{
		Matchmaker: &matchmaker,
		TaskCh:     &taskCh,
		FriendList: &friendList,
	}

	go listener.Listen(serviceStore)

	resolveTasks(serviceStore)
}

func resolveTasks(serviceStore ss.ServiceStore) {
	for {
		task, ok := <-*serviceStore.TaskCh

		var newTasks []dtasks.ITask

		if !ok {
			fmt.Println("Channel closed")
			return
		}

		switch task.GetTaskId() {
		case dtasks.AskInfoId:
			newTasks = resolveAskInfo(task.(*dtasks.AskInfoTask), serviceStore)
		case dtasks.SortInfoId:
			newTasks = resolveSortInfo(task.(*dtasks.SortInfoTask), serviceStore)
		default:
			fmt.Println("Uknown task")
		}

		for _, v := range newTasks {
			*serviceStore.TaskCh <- v
		}
	}
}

func resolveAskInfo(task *dtasks.AskInfoTask, serviceStore ss.ServiceStore) []dtasks.ITask {

	if task.Result == nil {
		if !serviceStore.Matchmaker.ProcessAskInfoTask(task) {
			return nil
		}
	}

	broadcaster.ProcessAskInfoTask(task)
	return nil
}

func resolveSortInfo(task *dtasks.SortInfoTask, serviceStore ss.ServiceStore) []dtasks.ITask {
	return convertArrayToITask(serviceStore.Matchmaker.ProcessSortInfoTask(task))
}
