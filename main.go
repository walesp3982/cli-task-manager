package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"task-manager-cli/cmd"
	"task-manager-cli/task"
)

const HELP_OPTION = "Comandos básicos: " +
	"\n - add <name> | Agregar tarea " +
	"\n - remove <id> | Eliminar tarea " +
	"\n - done <id> | Marcar tarea como realizada " +
	"\n - list | Listar tareas"

func main() {

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("No se encontró argumentos, por favor usar help para más información ")
		return
	}

	mode := arguments[0]
	if mode == "help" {
		fmt.Println(HELP_OPTION)
		return
	}

	dirActual, err := os.Getwd()
	if err != nil {
		fmt.Println("Cannot access to current directory")
		return
	}

	path := filepath.Join(dirActual, "tasks.json")

	if mode == "list" {
		tasks, err := cmd.List(path)
		if err != nil {
			fmt.Println("No se pudo carga el archivo, error")
			return
		}
		cmd.GenerateTable(tasks)
		return
	}

	if len(arguments) < 2 {
		fmt.Println("No se puso el argumento requerido, use help para más información")
		return
	}

	argument := arguments[1]

	if mode == "add" {
		cmd.AddTask(argument, path)
		fmt.Printf("Tarea < %s > agregada correctamente ", argument)
		return
	}

	if mode == "remove" {
		id, err := strconv.Atoi(argument)
		if err != nil {
			fmt.Println("Cannot convert id param to int")
			return
		}
		delete_task, err := cmd.DeleteTask(id, path)
		if err != nil {
			fmt.Println("La tarea no fue eliminada")
			return
		}
		fmt.Println("Se eliminó la tarea")
		cmd.GenerateTable([]task.Task{*delete_task})
		return
	}

	if mode == "done" {
		id, err := strconv.Atoi(argument)
		if err != nil {
			fmt.Println("Cannot convert id param to int")
			return
		}
		cmd.Done(id, path)
		fmt.Println("Marcando tarea como completada")
		return
	}

	fmt.Println("Argumentos inválido, use help para más información")
}
