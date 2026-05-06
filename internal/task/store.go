package task

import (
	"encoding/json"
	"os"
)

const taskFile = "tasks.json"

// LoadTasks membaca semua tugas dari file
func LoadTasks() ([]Task, error) {
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

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func FindIndex(tasks []Task, id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}
