package listener

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func Listen() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	err := serverOne.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
}
