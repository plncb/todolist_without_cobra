package cmd

import (
	"flag"
	"fmt"
	"log"
	"time"
	"todolist/internal/database"
)

func AddTask(args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	var dueBy string
	fs.StringVar(&dueBy, "date", "9999-12-31", "Date in YYYY-MM-DD format")

	fs.Parse(args)

	if len(fs.Args()) < 1 {
		log.Fatal("Task description is required.")
	}

	taskDescription := fs.Arg(0)
	const dateFormat = "2006-01-02" // Format: YYYY-MM-DD
	parsedDate, err := time.Parse(dateFormat, dueBy)
	if err != nil {
		log.Fatalf("Invalid date format. Please use YYYY-MM-DD.")
	}

	_, err = database.CreateTask(database.DB, taskDescription, parsedDate)
	if err != nil {
		log.Fatalf("Error creating task: %v", err)
	}

	fmt.Println("Task added successfully.")
}

func PrintAddTaskHelp() {
	fmt.Println(`Add a task

Usage:
  todolist add [flags]

Flags:
  -d, --date string   Date in YYYY-MM-DD format (default "9999-12-31")
  -h, --help          help for add`)
}
