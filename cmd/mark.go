package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/AzafaDev/task-tracker-go/internal/task"
)

func MarkTask(status string, args []string) {
	if len(args) == 0 {
		fmt.Printf("Use: task-cli mark-%s <id>\n", status)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Id must be number")
		return
	}

	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error in loading data:", err)
		return
	}

	idx := task.FindIndex(tasks, id)
	if idx == -1 {
		fmt.Printf("No task found with id [%d]\n", id)
		return
	}

	tasks[idx].Status = status
	tasks[idx].UpdatedAt = time.Now().Format(time.RFC3339)

	if err := task.SaveTasks(tasks); err != nil {
		fmt.Println("Error in saving data:", err)
		return
	}

	fmt.Printf("Task [%d] is now %s\n", id, status)
}
