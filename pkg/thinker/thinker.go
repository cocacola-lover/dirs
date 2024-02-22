package thinker

import (
	fl "dirs/pkg/friendList"
	"dirs/pkg/listener"
	"dirs/pkg/logger"
	m "dirs/pkg/matchmaker"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
)

func InitThinker(logger logger.Logger) {
	matchmaker := m.NewMatchmaker(logger)
	friendList := fl.NewFriendList(logger)
	taskCh := make(chan dtasks.ITask)

	serviceStore := ss.ServiceStore{
		Matchmaker: &matchmaker,
		TaskCh:     &taskCh,
		FriendList: &friendList,
		Logger:     &logger,
	}

	go listener.Listen(serviceStore)
	go listener.Serve(serviceStore)

	resolveTasks(serviceStore)
}

func resolveTasks(serviceStore ss.ServiceStore) {
	for {
		task, ok := <-*serviceStore.TaskCh

		var newTasks []dtasks.ITask

		if !ok {
			serviceStore.Logger.Warning.Println("Channel closed")
			return
		}

		switch task.GetTaskId() {
		case dtasks.AskInfoId:
			newTasks = resolveAskInfo(task.(*dtasks.AskInfoTask), serviceStore)
		case dtasks.SortInfoId:
			newTasks = resolveSortInfo(task.(*dtasks.SortInfoTask), serviceStore)
		case dtasks.DemandInfoId:
			resolveDemandInfo(task.(*dtasks.DemandInfoTask), serviceStore)
		default:
			serviceStore.Logger.Warning.Println("Uknown task")
		}

		for _, v := range newTasks {
			*serviceStore.TaskCh <- v
		}
	}
}
