package master

import (
	"fmt"
	"net/rpc"
)

func rpcCall(methodName string, arg any, reply any) error {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return err
	}
	defer client.Close()

	err = client.Call(methodName, struct{}{}, reply)
	if err != nil {
		msg := fmt.Sprintf("Error calling %s:", methodName)
		fmt.Println(msg, err)
		return err
	}

	return nil
}

func OrderGreet() (string, error) {
	var reply string

	err := rpcCall("Tools.Greet", struct{}{}, &reply)
	return reply, err
}
