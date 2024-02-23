package listener

import (
	"context"
	envp "dirs/pkg/environment"
	tp "dirs/pkg/tasks"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type butlerKeyType string

const envKeyButler butlerKeyType = "envButler"
const chKeyButler butlerKeyType = "chButler"

func receiveDemand(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", tp.NewDemandInfoTaskPointer, chKeyButler)
}

func receivePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! Everything is working as intended\n")
}

func receiveLogRequest(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(b))
}

func Serve(env envp.Environment, taskCh *chan tp.ITask) {
	messageContext := context.Background()

	// Add env
	messageContext = context.WithValue(messageContext, envKeyButler, env)
	// Add taskCh
	messageContext = context.WithValue(messageContext, chKeyButler, taskCh)

	mux := http.NewServeMux()
	mux.HandleFunc("/ask", receiveDemand)
	mux.HandleFunc("/ping", receivePing)
	mux.HandleFunc("/log", receiveLogRequest)
	serverOne := &http.Server{
		Addr:    ":3334",
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
