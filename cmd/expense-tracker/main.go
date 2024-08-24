package main

import (
	"expense-tracker/internal/expense"
	"expense-tracker/internal/storage"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := "expenses.json"
	store, err := storage.NewFileStorage(filePath)
	if err != nil {
		log.Fatal(err)
	}

	service := expense.NewExpenseService(store)

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0, "Expense amount")

	delCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := delCmd.Int("id", 0, "Expense ID to delete")

	sumCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := sumCmd.Int("month", 0, "Month number (1-12)")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'delete', 'list', or 'summary' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *desc == "" || *amount <= 0 {
			fmt.Println("Please provide a valid description and amount")
			os.Exit(1)
		}
		expense, err := service.AddExpense(*desc, *amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)

	case "delete":
		delCmd.Parse(os.Args[2:])
		if *delID == 0 {
			fmt.Println("Please provide a valid ID to delete")
			os.Exit(1)
		}
		if err := service.DeleteExpense(*delID); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Expense deleted successfully")
	case "list":
		expenses, err := service.ListExpenses()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID Date Description Amount")
		for _, exp := range expenses {
			fmt.Printf("%d   %s  %s  $%.2f\n", exp.ID, exp.Date.Format("2006-01-02"), exp.Description, exp.Amount)
		}
	case "summary":
		sumCmd.Parse(os.Args[2:])
		var total float64
		if *month == 0 {
			expenses, err := service.ListExpenses()
			if err != nil {
				log.Fatal(err)
			}
			for _, exp := range expenses {
				total += exp.Amount
			}
		} else {
			total, err = service.Summary(*month)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("Total expenses: $%.2f\n", total)
	default:
		fmt.Println("Expected 'add', 'delete', 'list', or 'summary' subcommands")
		os.Exit(1)
	}
}
