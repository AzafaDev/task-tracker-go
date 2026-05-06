package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/AzafaDev/task-tracker-go/internal/task"
)

func AddTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Use: task-cli add <title>")
		return
	}
	title := args[0]
	if strings.TrimSpace(title) == "" {
		fmt.Println("Please add title correctly!")
		return
	}

	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks:", err)
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	now := time.Now().Format(time.RFC3339)
	newTask := task.Task{
		ID:        newID,
		Title:     title,
		Status:    "todo",
		CreatedAt: now,
		UpdatedAt: now,
	}

	tasks = append(tasks, newTask)
	if err := task.SaveTasks(tasks); err != nil {
		fmt.Println("Failed to save:", err)
		return
	}
	fmt.Printf("Added task: [%d] %s\n", newID, title)
}
