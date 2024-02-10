package models

import "fmt"

type Expense struct {
	Name       string
	Amount     string
	Date       string // Should this be a Date type instead?
	Categories []string
}

func (exp Expense) String() string {
	return fmt.Sprintf("[Name: %s, Amount: %s, Date: %s, Categories: (%s)]", exp.Name, exp.Amount, exp.Date, exp.Categories)
}
