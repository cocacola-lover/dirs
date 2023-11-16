package main

import (
	"dirs/pkg/listener"
	dtasks "dirs/pkg/tasks"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	taskCh := make(chan dtasks.BaseTask)
	go listener.Listen(taskCh)

	for {
		task, ok := <-taskCh

		if !ok {
			fmt.Println("Channel closed")
			return
		}

		fmt.Printf("New task : %s\n", task.String())
	}
}
