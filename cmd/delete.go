package cmd

import (
	"errors"
	"task-manager-cli/task"
)

func DeleteTask(id int, path string) (*task.Task, error) {
	if !task.FileExist(path) {
		return nil, errors.New("File not found error")
	}

	tasks := task.GetTasks(path)

	delete_task := tasks.GetTaskById(id)

	if delete_task == nil {
		return nil, errors.New("Task not found")
	}

	tasks.DeleteTask(id)
	task.SaveTasks(tasks, path)
	return delete_task, nil
}
