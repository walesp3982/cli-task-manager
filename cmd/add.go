package cmd

import (
	"task-manager-cli/task"
)

func AddTask(name string, path string) {
	if !task.FileExist(path) {
		task.CreateStorage(path)
		tasks := task.CreateTasks()
		task.SaveTasks(*tasks, path)
	}
	// Get information on JSON
	tasks := task.GetTasks(path)
	// AddTask
	tasks.AddTask(name)

	// Writing in the file
	task.SaveTasks(tasks, path)

}
