package models

import "fmt"

type Expense struct {
	Name       string
	Amount     float32
	Date       string
	Categories []ExpenseCategory
	Note       string
}

func (exp Expense) String() string {
	return fmt.Sprintf("[Name: %s, Amount: %.2f, Date: %s, Categories: (%s), Note: (%s)]",
		exp.Name, exp.Amount, exp.Date, exp.Categories, exp.Note)
}
