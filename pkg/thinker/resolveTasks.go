package thinker

import (
	"dirs/pkg/broadcaster"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
)

func resolveAskInfo(task *tp.AskInfoTask, env envp.Environment) []tp.ITask {

	if task.Result == nil {
		if !env.Matchmaker.ProcessAskInfoTask(task) {
			return nil
		}
	}

	broadcaster.ProcessAskInfoTask(task)
	return nil
}

func resolveSortInfo(task *tp.SortInfoTask, env envp.Environment) []tp.ITask {
	return convertArrayToITask(env.Matchmaker.ProcessSortInfoTask(task))
}

func resolveDemandInfo(task *tp.DemandInfoTask, env envp.Environment) {
	broadcaster.ProcessDemandInfoTask(task, env)
}
