package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RAHULBONEY/clITaskTracker/internal/task"
)

func main() {
	var input string
	err := task.LoadTasks()
	if err != nil {
		fmt.Println("Warning: Could not load tasks:", err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a full command: ")
	if scanner.Scan() {
		input = scanner.Text()
		fmt.Println("You typed:", input)
	}
	parts := strings.SplitN(input, " ", 3)
	var command, taskName string

	if len(parts) > 1 {
		command = parts[1]
		fmt.Println("Command", command)
	}
	if len(parts) > 2 {
		taskName = parts[2]
		fmt.Println("Task", taskName)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading standard input:", err)
	}
	switch {
	case len(command) == 0:
		fmt.Println("Enter the commands from add,get,update,delete")

	case command == "add":
		if taskName == "" {
			fmt.Println("Error:Task creation error")
		} else {
			task.AddTask(taskName)
			fmt.Printf("Task created successfully '%s'\n", taskName)

		}
	case command == "update" || command == "complete":
		if taskName == "" {
			fmt.Println("Error: Please provide a task ID.")
		} else {

			taskID, err := strconv.Atoi(taskName)
			if err != nil {
				fmt.Println("Error: Task ID must be a valid number.")
			} else {

				err = task.CompleteTask(taskID)
				if err != nil {

					fmt.Println("Error:", err)
				} else {
					fmt.Printf("Task %d marked as completed.\n", taskID)
				}
			}
		}
	case command == "delete":
		if taskName == "" {
			fmt.Println("Error: Please provide a task ID.")
		} else {
			taskID, err := strconv.Atoi(taskName)
			if err != nil {
				fmt.Println("Error: Task ID must be a valid number.")
			} else {
				err = task.DeleteTask(taskID)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Printf("Task %d deleted successfully.\n", taskID)
				}
			}
		}
	case command == "get":
		fmt.Println("Fetched tasks")
		task.GetTask()
	default:
		fmt.Println("Command empty")
	}

}
