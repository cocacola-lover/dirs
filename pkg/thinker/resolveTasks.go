package thinker

import (
	"dirs/pkg/broadcaster"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
)

func resolveAskInfo(task *tp.AskInfoTask, env envp.Environment) []tp.ITask {

	if task.Result == nil {
		_, alreadyRegistered := env.Matchmaker.ProcessAskInfoTask(task)

		if alreadyRegistered {
			return nil
		}
	}

	broadcaster.ProcessAskInfoTask(task, env)
	return nil
}

func resolveSortInfo(task *tp.SortInfoTask, env envp.Environment) []tp.ITask {
	return convertArrayToITask(env.Matchmaker.ProcessSortInfoTask(task))
}

func resolveDemandInfo(task *tp.DemandInfoTask, env envp.Environment) {
	broadcaster.ProcessDemandInfoTask(task, env)
}
