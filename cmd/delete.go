package cmd

import (
	"fmt"
	"log"
	"strconv"
	"todolist/internal/database"
)

func DeleteTask(args []string) {
	if len(args) < 1 {
		log.Fatal("Task ID is required.")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Invalid task ID: %v", err)
	}

	err = database.DeleteTask(database.DB, id)
	if err != nil {
		log.Fatalf("Error deleting task: %v", err)
	}

	fmt.Printf("Task %d successfully deleted.\n", id)
}

func PrintDeleteTaskHelp() {
	fmt.Println(`Delete a task based on its id

Usage:
  todolist delete [id]

Flags:
  -h, --help   help for delete`)
}
