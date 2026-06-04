package cmd

import (
	"errors"
	"task-manager-cli/task"
)

func Done(id int, path string) error {
	if !task.FileExist(path) {
		return errors.New("File not found")
	}

	tasks := task.GetTasks(path)

	task_found := tasks.GetTaskById(id)

	if task_found == nil {
		return errors.New("Task not found")
	}

	task_found.Done = true

	task.SaveTasks(tasks, path)

	return nil
}
