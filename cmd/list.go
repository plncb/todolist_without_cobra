package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"todolist/internal/database"

	"github.com/mergestat/timediff"
)

func ListTasks(args []string) {
	tasks, err := database.GetAllTasks(database.DB)
	if err != nil {
		log.Fatalf("Error fetching tasks: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Task ID\tDescription\tCreated\tDue By\tCompleted")
	for _, task := range tasks {
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%t\n",
			task.ID,
			task.Description,
			timediff.TimeDiff(task.CreatedAt),
			timediff.TimeDiff(task.DueBy),
			task.IsCompleted,
		)
	}
	w.Flush()
}

func PrintListTasksHelp() {
	fmt.Println(`List all tasks

Usage:
  todolist list

Flags:
  -h, --help   help for list`)
}
