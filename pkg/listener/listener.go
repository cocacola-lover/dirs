package listener

import (
	"context"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type listenerKeyType string

const ssKeyListener listenerKeyType = "ssKeyListener"

func askForInfo(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", dtasks.NewAskInfoTaskPointer, ssKeyListener)
}

func receiveInfo(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", dtasks.NewSortInfoTaskPointer, ssKeyListener)
}

func Listen(serviceStore ss.ServiceStore) {
	messageContext := context.Background()
	messageContext = context.WithValue(messageContext, ssKeyListener, serviceStore)

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
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
}
