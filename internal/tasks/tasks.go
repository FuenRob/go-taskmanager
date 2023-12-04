package tasks

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          int
	Name        string
	Description string
	Completed   bool
}

func Add(db *gorm.DB, task Task) *Task {
	task.Completed = false
	db.Create(&task)
	return &task
}

func GetAll(db *gorm.DB) []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func GetByID(db *gorm.DB, id int) *Task {
	var task Task
	db.Find(&task, id)
	return &task
}

func DeleteByID(db *gorm.DB, id int) {
	result := db.Delete(&Task{}, id)
	if result.Error != nil {
		panic(result.Error)
	}
}

func UpdateByID(db *gorm.DB, id int, task Task) *Task {
	var taskToUpdate Task
	resultTask := db.First(&taskToUpdate, id)
	if resultTask.Error != nil {
		panic(resultTask.Error)
	}

	taskToUpdate = task

	resultSave := db.Save(&taskToUpdate)

	if resultSave.Error != nil {
		panic(resultSave.Error)
	}

	return &taskToUpdate
}
