package expense

import (
	"errors"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}

func NewExpense(id int, description string, amount float64) Expense {
	return Expense{
		ID:          id,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
}

type ExpenseManager struct {
	Expenses []Expense
	NextID   int
}

func NewExpenseManager() *ExpenseManager {
	return &ExpenseManager{
		Expenses: []Expense{},
		NextID:   1,
	}
}

func (em *ExpenseManager) Add(description string, amount float64) Expense {
	expense := NewExpense(em.NextID, description, amount)
	em.Expenses = append(em.Expenses, expense)
	em.NextID++
	return expense
}

func (em *ExpenseManager) Delete(id int) error {
	for i, expense := range em.Expenses {
		if expense.ID == id {
			em.Expenses = append(em.Expenses[:i], em.Expenses[i+1:]...)
			return nil
		}
	}
	return errors.New("expense not found")
}

func (em *ExpenseManager) GetAll() []Expense {
	return em.Expenses
}

func (em *ExpenseManager) Summary() float64 {
	var total float64
	for _, expense := range em.Expenses {
		total += expense.Amount
	}
	return total
}
