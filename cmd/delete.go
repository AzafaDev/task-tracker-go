package cmd

import (
	"fmt"
	"strconv"

	"github.com/AzafaDev/task-tracker-go/internal/task"
)

func DeleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Use: task-cli delete <id>")
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

	tasks = append(tasks[:idx], tasks[idx+1:]...)
	if err := task.SaveTasks(tasks); err != nil {
		fmt.Println("Error in saving data:", err)
		return
	}
	fmt.Printf("Task [%d] deleted successfully\n", id)
}
