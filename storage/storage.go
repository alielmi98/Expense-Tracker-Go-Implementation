package storage

import (
	"encoding/json"
	"os"

	"github.com/alielmi98/Expense-Tracker-Go-Implementation/expense"
)

func SaveExpenses(filename string, em *expense.ExpenseManager) error {
	data, err := json.Marshal(em)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadExpenses(filename string) (*expense.ExpenseManager, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return expense.NewExpenseManager(), nil // Return a new manager if the file doesn't exist
	}

	var em expense.ExpenseManager
	if err := json.Unmarshal(data, &em); err != nil {
		return nil, err
	}
	return &em, nil
}
