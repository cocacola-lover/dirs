package listener

import (
	"context"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
)

type keyType string

const envKey keyType = "env"
const chKey keyType = "ch"

func extractValues(sc context.Context) (envp.Environment, chan tp.ITask) {
	env := sc.Value(envKey).(envp.Environment)
	taskCh := *sc.Value(chKey).(*chan tp.ITask)

	return env, taskCh
}

func newServeContext(env envp.Environment, taskCh *chan tp.ITask) context.Context {
	sc := context.Background()

	sc = context.WithValue(sc, envKey, env)
	sc = context.WithValue(sc, chKey, taskCh)

	return sc
}
