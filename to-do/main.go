package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Simple ToDo App")

	tasks := []string{}

	for {
		fmt.Println("Enter a task or press 'q' to quit")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "q" {
			fmt.Println("Quitting the application")
			break
		}
		tasks = append(tasks, input)
		fmt.Printf("Task added %s \n", input)
	}

	fmt.Println("\nToDo List:")
	for index, task := range tasks {
		fmt.Printf("%d %s \n", index+1, task)
	}
}
