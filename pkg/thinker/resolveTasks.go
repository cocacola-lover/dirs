package thinker

import (
	"dirs/pkg/broadcaster"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
)

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