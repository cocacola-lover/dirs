package listener

import (
	"context"
	drequests "dirs/pkg/requests"
	ss "dirs/pkg/serviceStore"
	dtasks "dirs/pkg/tasks"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

type keyType string

const ssKey keyType = "ssKey"

func askForInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var askInfoR drequests.AskInfoRequest
	marshalErr := json.Unmarshal(body, &askInfoR)
	if marshalErr != nil {
		http.Error(w, "Wrong json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	taskCh := *ctx.Value(ssKey).(ss.ServiceStore).TaskCh
	taskCh <- dtasks.NewOuterAskInfoTaskPointer(askInfoR)

	fmt.Printf("got /ask request %s\n", string(body))
}

func receiveInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var sendInfoR drequests.SendInfoRequest
	marshalErr := json.Unmarshal(body, &sendInfoR)
	if marshalErr != nil {
		http.Error(w, "Wrong json", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	taskCh := *ctx.Value(ssKey).(ss.ServiceStore).TaskCh
	taskCh <- dtasks.NewSortInfoTaskPointer(sendInfoR)

	fmt.Printf("got /send request %s\n", string(body))
}

func Listen(serviceStore ss.ServiceStore) {
	messageContext := context.Background()
	messageContext = context.WithValue(messageContext, ssKey, serviceStore)

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
