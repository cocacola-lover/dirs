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

func receiveDemand(w http.ResponseWriter, r *http.Request) {
	env, taskCh := extractValues(r.Context())

	newTask, err := readRequestAndCreateTask(w, r, http.MethodPost, tp.NewDemandInfoTaskPointer)
	if err != nil {
		env.Error.Printf("Error while receiving demand : %s\n", err)
		return
	}

	env.Info.Printf("Received demand : %s", *newTask)
	taskCh <- *newTask
}

func receivePing(w http.ResponseWriter, r *http.Request) {
	env, _ := extractValues(r.Context())
	env.Info.Println("Received ping request")

	fmt.Fprintf(w, "Hello! Everything is working as intended\n")
}

func receiveLogRequest(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("logs.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(b))

	env, _ := extractValues(r.Context())
	env.Info.Println("Received log request")
}

func Serve(env envp.Environment, taskCh *chan tp.ITask) {
	mux := http.NewServeMux()
	mux.HandleFunc("/ask", receiveDemand)
	mux.HandleFunc("/ping", receivePing)
	mux.HandleFunc("/log", receiveLogRequest)
	serverOne := &http.Server{
		Addr:    ":3334",
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
