package cmd

import (
	"fmt"

	"github.com/AzafaDev/task-tracker-go/internal/task"
)

func ListTasks(args []string) {
	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load data:", err)
		return
	}

	filter := ""
	if len(args) >= 1 {
		filter = args[0]
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks provided")
		return
	}

	fmt.Println("ID  | Status        | Judul")
	fmt.Println("----|---------------|------")
	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}
		fmt.Printf("%-4d| %-13s | %s\n", t.ID, t.Status, t.Title)
	}
}
