package cmd

import (
	"fmt"
	"os"
	"strings"
	"taskmanager/internal/tasks"
	"taskmanager/utils"

	"github.com/spf13/cobra"
)

func Execute() {

	var cmdAdd = &cobra.Command{
		Use:   "add [string]",
		Short: "Command to add a task",
		Long:  `This command add a task to the task list.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var task = tasks.Task{
				Name: strings.Join(args, " "),
			}
			newTask := tasks.Add(task)

			fmt.Printf(`Tarea creada con éxito: %s(%d)`, newTask.Name, task.ID)

		},
	}

	var cmdList = &cobra.Command{
		Use:   "list",
		Short: "Command to list all tasks",
		Long:  `This command list all tasks.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(tasks.GetAll())

		},
	}

	var cmdDetail = &cobra.Command{
		Use:   "detail [id]",
		Short: "Command to show a task detail",
		Long:  `This command show a task detail.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])
			fmt.Println(tasks.GetByID(id))

		},
	}

	var cmdUpdate = &cobra.Command{
		Use:   "update [id] [string] [string]",
		Short: "Command to update a task",
		Long:  `This command update a task.`,
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			var tasksEdited = tasks.GetByID(id)

			if args[1] == "name" {
				tasksEdited.Name = args[2]
			} else if args[1] == "description" {
				tasksEdited.Description = args[2]
			} else {
				panic("Invalid argument")
			}

			fmt.Println(tasks.UpdateByID(id, *tasksEdited))

		},
	}

	var cmdCompleted = &cobra.Command{
		Use:   "completed [id]",
		Short: "Command to mark a task as completed",
		Long:  `This command mark a task as completed.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			var tasksEdited = tasks.GetByID(id)

			tasksEdited.Completed = true

			fmt.Println(tasks.UpdateByID(id, *tasksEdited))
		},
	}

	var cmdDeleted = &cobra.Command{
		Use:   "delete [id]",
		Short: "Command to delete a task",
		Long:  `This command delete a task.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := utils.ParseInt(args[0])

			tasks.DeleteByID(id)
		},
	}

	var rootCmd = &cobra.Command{Use: "taskmanager"}
	rootCmd.AddCommand(cmdAdd)
	rootCmd.AddCommand(cmdList)
	rootCmd.AddCommand(cmdDetail)
	rootCmd.AddCommand(cmdUpdate)
	rootCmd.AddCommand(cmdCompleted)
	rootCmd.AddCommand(cmdDeleted)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
