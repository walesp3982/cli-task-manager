package cmd

import (
	"errors"
	"fmt"
	"task-manager-cli/task"
)

func List(path string) ([]task.Task, error) {
	if !task.FileExist(path) {
		return nil, errors.New("File not found")
	}

	tasks := task.GetTasks(path)

	return tasks.Tasks, nil
}

func GenerateTable(tasks []task.Task) {
	if len(tasks) == 0 {
		fmt.Println(" -> No se encontró ninguna tarea")
		fmt.Println(" (Por favor ingresar alguna tarea)")
	}
	for _, task := range tasks {
		str_done := "Not done"
		if task.Done {
			str_done = "Done"
		}
		fmt.Printf("> | %d | %s | %s | \n", task.Id, task.Name, str_done)
	}
}
