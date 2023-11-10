package butler

import (
	"fmt"
	"net"
	"net/rpc"
)

func InitButler() {
	butler := new(Tools)
	rpc.Register(butler)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	fmt.Println("Butler listening on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
