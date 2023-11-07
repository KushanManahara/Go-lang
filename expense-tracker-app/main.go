package main

import (
	"bufio"
	"fmt"
	"os"
)

type Expense struct {
	Date    string
	Details string
	Amount  string
}

var expenses []Expense

func main() {

	fmt.Println("Expense Tracker")

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Exit")
		fmt.Print("Enter option: ")

		var choise int

		fmt.Println("Enter your choise: ")
		fmt.Scan(&choise)

		switch choise {
		case 1:
			recordExpense()

		case 2:
			viewExpense()

		case 3:
			fmt.Println("Exiting Expense Tracker")
			os.Exit(0)
		default:
			fmt.Println("Invalid choise, Please try again")

		}
	}
}

func recordExpense() {
	var expense Expense

	fmt.Print("Enter date (e.g., 2023-11-01): ")
	fmt.Scan(&expense.Date)

	fmt.Print("Enter expense details: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expense.Details = scanner.Text()

	fmt.Print("Enter expense amount: ")
	fmt.Scan(&expense.Amount)

	expenses = append(expenses, expense)
	fmt.Println("Expense recorded successfully!")
}

func viewExpense() {
	fmt.Println("\nExpense List:")
	for _, expense := range expenses {
		fmt.Printf("Date: %s\nDetails: %s\nAmount: $%.2f\n\n", expense.Date, expense.Details, expense.Amount)
	}
}
