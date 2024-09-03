package cmd

import "fmt"

func PrintHelp() {
	fmt.Println(`Usage:
  todolist [command]

Available Commands:
  add         Add a task
  complete    Complete a task based on its id
  delete      Delete a task based on its id
  help        Help about any command
  list        List all tasks

Flags:
  -h, --help     help for todolist
  -t, --toggle   Help message for toggle

Use "todolist [command] --help" for more information about a command.`)
}
