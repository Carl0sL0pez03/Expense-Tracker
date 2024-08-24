package expense

import (
	"expense-tracker/internal/storage"
	"expense-tracker/pkg/model"
	"time"
)

type ExpenseService struct {
	storage storage.Storage
}

func NewExpenseService(storage storage.Storage) *ExpenseService {
	return &ExpenseService{storage: storage}
}

func (s *ExpenseService) AddExpense(description string, amount float64) (*model.Expense, error) {
	expense := &model.Expense{
		ID:          s.storage.GetNextID(),
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
	if err := s.storage.SaveExpense(expense); err != nil {
		return nil, err
	}
	return expense, nil
}

func (s *ExpenseService) DeleteExpense(id int) error {
	return s.storage.DeleteExpense(id)
}

func (s *ExpenseService) ListExpenses() ([]*model.Expense, error) {
	return s.storage.ListExpenses()
}

func (s *ExpenseService) Summary(month int) (float64, error) {
	expenses, err := s.storage.ListExpenses()
	if err != nil {
		return 0, err
	}

	var total float64
	for _, exp := range expenses {
		if exp.Date.Month() == time.Month(month) {
			total += exp.Amount
		}
	}

	return total, nil
}
