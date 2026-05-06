package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AzafaDev/task-tracker-go/internal/task"
)

func UpdateTask(args []string) {
	if len(args) < 2 {
		fmt.Println("Use: task-cli update <id> <new title>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Id must be number")
		return
	}

	newTitle := args[1]
	if newTitle == "" {
		fmt.Println("Please provide new title correctly!")
		return
	}

	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks:", err)
		return
	}

	idx := task.FindIndex(tasks, id)
	if idx == -1 {
		fmt.Printf("Task with id [%d] not found\n", id)
		return
	}

	tasks[idx].Title = newTitle
	tasks[idx].UpdatedAt = time.Now().Format(time.RFC3339)

	if err := task.SaveTasks(tasks); err != nil {
		fmt.Println("Error in saving task:", err)
		return
	}
	fmt.Printf("Task [%d] updated successfully\n", id)
}
