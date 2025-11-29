package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func add(tasks []string) []string {
	fmt.Println("Enter task:")
	task := readLine()
	tasks = append(tasks, task)
	fmt.Println("New Task added:", task)
	return tasks
}

func list(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func deleteTask(tasks []string) []string {
	fmt.Println("Enter task number to delete:")
	input := readLine()

	id := 0
	fmt.Sscanf(input, "%d", &id)

	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid task ID")
		return tasks
	}

	tasks = append(tasks[:id-1], tasks[id:]...)
	fmt.Println("Task deleted successfully.")
	return tasks
}

func CompleteTask(tasks []string, id int) []string {
	if id < 1 || id > len(tasks) {
		fmt.Println("Invalid task ID")
		return tasks
	}

	tasks[id-1] = "✔ " + tasks[id-1]
	return tasks
}


func clear(tasks []string) []string {
	return []string{}
}

func printHelp() {
	fmt.Println(`
Available Commands:

  add           - Add a new task
  list          - Show all tasks
  delete        - Delete a task
  clear         - Remove all tasks
  help          - Show this help menu
  exit          - Exit the application
`);
}

func main() {
	fmt.Println("WELCOME TO THE TODO-APP")
	fmt.Println("Type 'help' to know commands.\n")

	var tasks []string

	for {
		fmt.Print("\nEnter command: ")
		command := readLine()

		switch command {
		case "add":
			tasks = add(tasks)

		case "list":
			list(tasks)

		case "delete":
			tasks = deleteTask(tasks)

		case "clear":
			tasks = clear(tasks)
			fmt.Println("All tasks cleared.")

		case "help":
			printHelp()

		case "exit":
			fmt.Println("Goodbye!")
			return
		
		case "done":
				fmt.Print("Enter task number to mark as done: ")
				idStr := readLine()

				var id int
				fmt.Sscanf(idStr, "%d", &id)

				tasks = CompleteTask(tasks, id)


		default:
			fmt.Println("Invalid command. Type 'help' to see options.")
		}
	}
}
