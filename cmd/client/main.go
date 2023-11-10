package main

import (
	"dirs/pkg/master"
	"fmt"
)

func main() {
	ans, err := master.OrderGreet()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
