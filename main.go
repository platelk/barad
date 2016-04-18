package main

import (
	"barad/cmd"
	"barad/cmd/handlers"
	"fmt"
)

func main() {
	fmt.Println("Barad.")
	command := cmd.CreateCmd("ls")
	handler := handlers.CreateUnixHandler()

	if handler.IsHandling(command) {
		fmt.Println(command.Command, " is handle by ", handler.Name)
		cmd := handler.Handle(command)
		cmd.Exec()
	}
}
