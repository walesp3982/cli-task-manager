package task

import (
	"encoding/json"
	"os"
)

func CreateStorage(path string) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTaskJson(content []byte) TasksJson {
	task_json := TasksJson{}
	err := json.Unmarshal(content, &task_json)
	check(err)
	return task_json
}

func getBitsTaskJson(task_json TasksJson) []byte {
	buffer, err := json.MarshalIndent(task_json, "", "  ")
	check(err)
	return buffer
}

// Get task by file
func GetTasks(path string) TasksJson {
	buff := getBufferFile(path)
	tasks := getTaskJson(buff)
	return tasks
}

// Save task into the file
func SaveTasks(tasks TasksJson, path string) {
	buff := getBitsTaskJson(tasks)
	writeBufferInFile(path, buff)
}
func getBufferFile(path string) []byte {
	dat, err := os.ReadFile(path)
	check(err)
	return dat
}

func writeBufferInFile(path string, buffer []byte) {
	err := os.WriteFile(path, buffer, 0644)
	check(err)
}
