package weblogger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type websocketMessenger struct {
	channel chan byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wm websocketMessenger) broadcast(rw http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(rw, r, nil)
	log.Println("Starting connection")
	if err != nil {
		log.Print("Upgrade websocket connection:", err)
		return
	}
	defer conn.Close()
	for {
		w, err := conn.NextWriter(websocket.TextMessage)
		if err != nil {
			log.Println("Create WebsocketWriter :", err)
			continue
		}

		if _, err := io.Copy(w, wm); err != nil {
			log.Println("Copy info to WebsocketWriter :", err)
			continue
		}
		if err := w.Close(); err != nil {
			log.Println("Close WebsocketWriter :", err)
			continue
		}
	}
}

func (wm websocketMessenger) broadcastLogs() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", wm.broadcast)

	server := &http.Server{
		Addr:    "0.0.0.0:3335",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
}

// This reader awaits info if channel is empty at the start
func (wm websocketMessenger) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	checkValue, ok := <-wm.channel

	if !ok {
		return 0, io.EOF
	}

	p[0] = checkValue

	for i := 1; i < len(p); i++ {
		select {
		case p[i] = <-wm.channel:
			// Do nothing
		default:
			return i, io.EOF
		}
	}

	return 0, nil
}

func (wm websocketMessenger) Write(p []byte) (n int, err error) {
	for _, val := range p {
		select {
		case wm.channel <- val:
			// Do nothing
		default:
			<-wm.channel
			wm.channel <- val
		}
	}

	return 0, nil
}

func newWebsocketMessenger() websocketMessenger {
	return websocketMessenger{channel: make(chan byte, 1024)}
}
