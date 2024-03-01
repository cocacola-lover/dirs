package listener

import (
	"context"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
	"errors"
	"net"
	"net/http"
)

func askForInfo(w http.ResponseWriter, r *http.Request) {
	env, taskCh := extractValues(r.Context())

	newTask, err := readRequestAndCreateTask(w, r, http.MethodPost, tp.NewAskInfoTaskPointer)
	if err != nil {
		env.Error.Printf("Error while receiving askRequest : %s\n", err)
		return
	}

	env.Info.Printf("Received askRequest : %s", *newTask)
	taskCh <- *newTask
}

func receiveInfo(w http.ResponseWriter, r *http.Request) {
	env, taskCh := extractValues(r.Context())

	newTask, err := readRequestAndCreateTask(w, r, http.MethodPost, tp.NewSortInfoTaskPointer)
	if err != nil {
		env.Error.Printf("Error while receiving askRequest : %s\n", err)
		return
	}

	env.Info.Printf("Received askRequest : %s", *newTask)
	taskCh <- *newTask
}

func Listen(env envp.Environment, taskCh *chan tp.ITask) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ask", askForInfo)
	mux.HandleFunc("/send", receiveInfo)
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return newServeContext(env, taskCh)
		},
	}

	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		env.Error.Printf("server one closed\n")
	} else if err != nil {
		env.Error.Printf("error listening for server one: %s\n", err)
	}
}
