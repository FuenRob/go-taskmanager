package main

import (
	"taskmanager/cmd"
	"taskmanager/internal/tasks"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "taskmanager_user:taskmanager_user_password@tcp(127.0.0.1:13306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&tasks.Task{})

	cmd.Execute(db)
}
