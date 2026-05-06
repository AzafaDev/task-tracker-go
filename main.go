package main

import (
	"fmt"
	"os"

	"github.com/AzafaDev/task-tracker-go/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command] [arguments...]")
		fmt.Println("Commands: add, list, update, delete, mark-in-progress, mark-done")
		os.Exit(1)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		cmd.AddTask(args)
	case "list":
		cmd.ListTasks(args)
	case "update":
		cmd.UpdateTask(args)
	case "delete":
		cmd.DeleteTask(args)
	case "mark-in-progress":
		cmd.MarkTask("in-progress", args)
	case "mark-done":
		cmd.MarkTask("done", args)
	default:
		fmt.Printf("Command %s is not recognized\n", command)
	}
}
