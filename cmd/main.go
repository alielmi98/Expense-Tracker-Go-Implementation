package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/alielmi98/Expense-Tracker-Go-Implementation/storage"
)

func main() {
	dataFile := filepath.Join("..", "data", "expenses.json")

	em, err := storage.LoadExpenses(dataFile)
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "Description of the expense")
	amount := addCmd.Float64("amount", 0, "Amount of the expense")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "ID of the expense to delete")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'delete' or 'list' or 'summary'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *description == "" || *amount <= 0 {
			fmt.Println("Please provide a valid description and amount.")
			return
		}
		expense := em.Add(*description, *amount)
		fmt.Printf("Expense added successfully (ID: %d)\n", expense.ID)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if err := em.Delete(*id); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Expense deleted successfully")
		}

	case "list":
		listCmd.Parse(os.Args[2:])
		expenses := em.GetAll()
		fmt.Println("ID  Date       Description  Amount")
		for _, expense := range expenses {
			fmt.Printf("%d   %s  %s  $%.2f\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
		}

	case "summary":
		summaryCmd.Parse(os.Args[2:])
		total := em.Summary()
		fmt.Printf("Total expenses: $%.2f\n", total)

	default:
		fmt.Println("Unknown command.")
		os.Exit(1)
	}

	if err := storage.SaveExpenses(dataFile, em); err != nil {
		fmt.Println("Error saving expenses:", err)
	}
}
