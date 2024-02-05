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

type butlerKeyType string

const ssKeyButler butlerKeyType = "ssKeyButler"

func receiveDemand(w http.ResponseWriter, r *http.Request) {
	readRequestAndCreateTask(w, r, "POST", dtasks.NewDemandInfoTaskPointer, ssKeyButler)
}

func Serve(serviceStore ss.ServiceStore) {
	messageContext := context.Background()
	messageContext = context.WithValue(messageContext, ssKeyButler, serviceStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/ask", receiveDemand)
	serverOne := &http.Server{
		Addr:    ":3334",
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