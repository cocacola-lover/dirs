package listener

import (
	"context"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
	"errors"
	"net"
	"net/http"
)

type listenerKeyType string

const envKeyListener listenerKeyType = "envListener"
const chKeyListener listenerKeyType = "chListener"

func askForInfo(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", tp.NewAskInfoTaskPointer, chKeyListener)
}

func receiveInfo(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", tp.NewSortInfoTaskPointer, chKeyListener)
}

func Listen(env envp.Environment, taskCh *chan tp.ITask) {
	messageContext := context.Background()

	// Add env
	messageContext = context.WithValue(messageContext, envKeyListener, env)
	// Add taskCh
	messageContext = context.WithValue(messageContext, chKeyListener, taskCh)

	mux := http.NewServeMux()
	mux.HandleFunc("/ask", askForInfo)
	mux.HandleFunc("/send", receiveInfo)
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return messageContext
		},
	}

	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		env.Error.Printf("server one closed\n")
	} else if err != nil {
		env.Error.Printf("error listening for server one: %s\n", err)
	}
}
