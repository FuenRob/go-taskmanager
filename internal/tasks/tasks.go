package tasks

type Task struct {
	ID          int
	Name        string
	Description string
	Completed   bool
}

var tasks = []*Task{}

func Add(task Task) *Task {
	task.ID = len(tasks) + 1
	task.Completed = false
	tasks = append(tasks, &task)
	return &task
}

func GetAll() []*Task {
	return tasks
}

func GetByID(id int) *Task {
	for _, task := range tasks {
		if task.ID == id {
			return task
		}
	}
	return nil
}

func DeleteByID(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
}

func UpdateByID(id int, task Task) *Task {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i] = &task
			return tasks[i]
		}
	}
	return nil
}
