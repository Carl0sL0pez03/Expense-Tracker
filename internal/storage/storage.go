package storage

import "expense-tracker/pkg/model"

type Storage interface {
	GetNextID() int
	SaveExpense(expense *model.Expense) error
	DeleteExpense(id int) error
	ListExpenses() ([]*model.Expense, error)
}
