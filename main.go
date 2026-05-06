package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

const taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(taskFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func findTaskIndex(tasks []Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command] [arguments...]")
		fmt.Println("Commands: add, list, update, delete, mark-in-progress, mark-done")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Use: task-cli add <title>")
			return
		}
		title := os.Args[2]
		if strings.TrimSpace(title) == "" {
			fmt.Println("Please add title correctly!")
			return
		}

		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Failed to load tasks")
			return
		}
		newId := 1
		if len(tasks) > 0 {
			newId = tasks[len(tasks)-1].ID + 1
		}
		newTask := Task{
			ID:        newId,
			Title:     title,
			Status:    "todo",
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
		tasks = append(tasks, newTask)
		if err := saveTask(tasks); err != nil {
			fmt.Println("Failed to save:", err)
			return
		}
		fmt.Printf("Added task: [%d] %s\n", newId, title)

	case "list":
		data, err := loadTasks()
		if err != nil {
			fmt.Println("Failed to load data:", err)
			return
		}
		filter := ""
		if len(os.Args) >= 3 {
			filter = os.Args[2]
		}
		if len(data) == 0 {
			fmt.Println("No tasks provided")
			return
		}
		result := []Task{}
		for _, task := range data {
			if filter != "" && task.Status != filter {
				continue
			}
			result = append(result, task)
		}
		fmt.Println("ID  | Status        | Judul")
		fmt.Println("----|---------------|------")
		for _, t := range result {
			fmt.Printf("%-4d| %-13s | %s\n", t.ID, t.Status, t.Title)
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Use: task-cli update <id> <new title>")
			return
		}
		if os.Args[3] == "" {
			fmt.Println("Please provide new title correctly!")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be number")
			return
		}
		newTitle := os.Args[3]
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("No tasks provided")
			return
		}
		index := findTaskIndex(tasks, id)
		if index == -1 {
			fmt.Printf("Task with id [%d] not found\n", id)
			return
		}
		tasks[index].Title = newTitle
		tasks[index].UpdatedAt = time.Now().Format(time.RFC3339)
		if err := saveTask(tasks); err != nil {
			fmt.Println("Error in saving task:", err)
			return
		}
		fmt.Printf("Tugas [%d] berhasil diperbarui.\n", id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Use: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be number")
			return
		}
		data, err := loadTasks()
		if err != nil {
			fmt.Println("Error in loading data:", err)
			return
		}
		idx := findTaskIndex(data, id)
		if idx == -1 {
			fmt.Printf("No task found with id [%d]\n", id)
			return
		}
		data = append(data[:idx], data[idx+1:]...)
		if err := saveTask(data); err != nil {
			fmt.Println("Error in saving data:", err)
			return
		}
		fmt.Printf("Task [%d] deleted successfully\n", id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Use: task-cli mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be number")
			return
		}
		data, err := loadTasks()
		if err != nil {
			fmt.Println("Error in loading data:", err)
			return
		}
		idx := findTaskIndex(data, id)
		if idx == -1 {
			fmt.Printf("No task found with id [%d]\n", id)
			return
		}
		data[idx].Status = "in-progress"
		data[idx].UpdatedAt = time.Now().Format(time.RFC3339)
		if err := saveTask(data); err != nil {
			fmt.Println("Error in saving data:", err)
			return
		}
		fmt.Printf("Task [%d] now is in-progress\n", id)

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Use: task-cli mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Id must be number")
			return
		}
		data, err := loadTasks()
		if err != nil {
			fmt.Println("Error in loading data:", err)
			return
		}
		idx := findTaskIndex(data, id)
		if idx == -1 {
			fmt.Printf("No task found with id [%d]\n", id)
			return
		}
		data[idx].Status = "done"
		data[idx].UpdatedAt = time.Now().Format(time.RFC3339)
		if err := saveTask(data); err != nil {
			fmt.Println("Error in saving data:", err)
			return
		}
		fmt.Printf("Task [%d] is now done\n", id)

	default:
		fmt.Printf("Command %s is not recognized\n", command)
	}
}
