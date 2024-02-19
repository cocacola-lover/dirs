package weblogger

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type websocketMessenger struct {
	channel chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  0,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (wm websocketMessenger) broadcast(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	log.Println("Starting connection")
	if err != nil {
		log.Print("Upgrade websocket connection:", err)
		return
	}
	defer conn.Close()
	for {
		if err := conn.WriteMessage(websocket.TextMessage, wm.ReadMessage()); err != nil {
			log.Println("Write websocket message :", err)
			break
		}
	}
	log.Println("Closing connection")
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
func (wm websocketMessenger) ReadMessage() []byte {
	return <-wm.channel
}

func (wm websocketMessenger) Write(p []byte) (n int, err error) {
	newp := make([]byte, len(p))
	copy(newp, p)

	select {
	case wm.channel <- newp:
		// Do nothing
	default:
		<-wm.channel
		wm.channel <- newp
	}

	return 0, nil
}

func newWebsocketMessenger() websocketMessenger {
	return websocketMessenger{channel: make(chan []byte, 16)}
}
