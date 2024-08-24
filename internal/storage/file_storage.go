package storage

import (
	"encoding/json"
	"errors"
	"expense-tracker/pkg/model"
	"io/ioutil"
	"os"
)

type FileStorage struct {
	filePath string
	expenses []*model.Expense
}

func NewFileStorage(filePath string) (*FileStorage, error) {
	fs := &FileStorage{filePath: filePath}
	if err := fs.load(); err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *FileStorage) load() error {
	file, err := os.Open(fs.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fs.expenses); err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) save() error {
	data, err := json.MarshalIndent(fs.expenses, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fs.filePath, data, 0644)
}

func (fs *FileStorage) GetNextID() int {
	if len(fs.expenses) == 0 {
		return 1
	}
	return fs.expenses[len(fs.expenses)-1].ID + 1
}

func (fs *FileStorage) SaveExpense(expense *model.Expense) error {
	fs.expenses = append(fs.expenses, expense)
	return fs.save()
}

func (fs *FileStorage) DeleteExpense(id int) error {
	for i, exp := range fs.expenses {
		if exp.ID == id {
			fs.expenses = append(fs.expenses[:i], fs.expenses[i+1:]...)
			return fs.save()
		}
	}
	return errors.New("expense not found")
}

func (fs *FileStorage) ListExpenses() ([]*model.Expense, error) {
	return fs.expenses, nil
}
