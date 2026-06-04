package task

type Task struct {
	Id   int
	Name string
	Done bool
}

func newTask(id int, name string) *Task {
	return &Task{
		Id:   id,
		Name: name,
		Done: false,
	}
}

type TasksJson struct {
	Tasks  []Task
	NextId int
}

func CreateTasks() *TasksJson {
	return &TasksJson{
		Tasks:  []Task{},
		NextId: 1,
	}
}

func Filter[T any](slice []T, test func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if test(item) {
			result = append(result, item)
		}
	}
	return result
}

// Create a new task
func (s *TasksJson) AddTask(name string) {
	new_task := newTask(s.NextId, name)
	s.Tasks = append(s.Tasks, *new_task)
	s.NextId += 1
}

func (s *TasksJson) DeleteTask(id int) {
	newTasks := Filter(s.Tasks, func(t Task) bool {
		return t.Id != id
	})

	s.Tasks = newTasks
}

func (s *TasksJson) GetTaskById(id int) *Task {
	for i, task := range s.Tasks {
		if task.Id == id {
			return &s.Tasks[i]
		}
	}
	return nil
}
