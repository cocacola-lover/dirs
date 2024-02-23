package thinker

import (
	envp "dirs/pkg/environment"
	"dirs/pkg/listener"
	tp "dirs/pkg/tasks"
)

func InitThinker(env envp.Environment) {

	taskCh := make(chan tp.ITask)

	go listener.Listen(env, &taskCh)
	go listener.Serve(env, &taskCh)

	resolveTasks(env, &taskCh)
}

func resolveTasks(env envp.Environment, taskCh *chan tp.ITask) {
	for {
		task, ok := <-*taskCh

		var newTasks []tp.ITask

		if !ok {
			env.Warning.Println("Channel closed")
			return
		}

		switch task.GetTaskId() {
		case tp.AskInfoId:
			newTasks = resolveAskInfo(task.(*tp.AskInfoTask), env)
		case tp.SortInfoId:
			newTasks = resolveSortInfo(task.(*tp.SortInfoTask), env)
		case tp.DemandInfoId:
			resolveDemandInfo(task.(*tp.DemandInfoTask), env)
		default:
			env.Warning.Println("Uknown task")
		}

		for _, v := range newTasks {
			*taskCh <- v
		}
	}
}
