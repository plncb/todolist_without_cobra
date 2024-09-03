package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todolist/internal/database"
)

func CompleteTask(args []string) {
	if len(args) < 1 {
		log.Fatal("Task ID is required.")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid task ID: %v", err)
	}

	err = database.CompleteTask(database.DB, id)
	if err != nil {
		log.Fatalf("Error completing task: %v", err)
	}

	fmt.Printf("Task %d marked as complete.\n", id)
}

func PrintCompleteTaskHelp() {
	fmt.Println(`Complete a task based on its id

Usage:
  todolist complete [id]

Flags:
  -h, --help   help for complete`)
}
