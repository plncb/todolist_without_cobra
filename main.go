package main

import (
	"os"
	"todolist/cmd"
	"todolist/internal/database"
)

func main() {
	database.Init("./internal/database/task.sqlite3")

	// Check if no command is provided or if help is requested.
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "-h" || os.Args[1] == "--help" {
		// Handle 'help' command for specific subcommands
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "add":
				cmd.PrintAddTaskHelp()
			case "complete":
				cmd.PrintCompleteTaskHelp()
			case "delete":
				cmd.PrintDeleteTaskHelp()
			case "list":
				cmd.PrintListTasksHelp()
			default:
				cmd.PrintHelp()
			}
		} else {
			cmd.PrintHelp()
		}
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		cmd.AddTask(os.Args[2:])
	case "complete":
		cmd.CompleteTask(os.Args[2:])
	case "delete":
		cmd.DeleteTask(os.Args[2:])
	case "list":
		cmd.ListTasks(os.Args[2:])
	default:
		cmd.PrintHelp()
		os.Exit(1)
	}
}
