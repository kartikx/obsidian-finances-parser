// Contains methods to generate and write data for Obsidian.
package obsidian

import (
	"fmt"
	"os"

	"github.com/kartikx/obsidian-finances-parser/pkg/models"
)

func FormatExpensesForObsidian(expenses []*models.Expense) ([]string, error) {
	formattedExpenses := make([]string, 0, len(expenses))

	for _, expense := range expenses {
		formattedExpense := formatExpenseForObsidian(expense)
		formattedExpenses = append(formattedExpenses, formattedExpense)
	}

	return formattedExpenses, nil
}

func formatExpenseForObsidian(expense *models.Expense) string {
	categories := ""

	for _, category := range expense.Categories {
		categories += category.String() + ","
	}

	return fmt.Sprintf("- #expense (name::%s) (amount::%.2f) (date::%s) (categories::%s) (note::%s)",
		expense.Name,
		expense.Amount,
		expense.Date,
		categories,
		expense.Note)
}

// TODO Rename this function and file.
func WriteToObsidianVault(formattedExpenses []string, outputFilePath string) error {
	fmt.Println("Writing to Obsidian Vault")

	if len(outputFilePath) == 0 {
		return fmt.Errorf("output file path provided is empty")
	}

	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	// TODO Appends should not write expenses that are already present. Will need to seek to figure this out.
	for _, formattedExpense := range formattedExpenses {
		_, err = file.WriteString(formattedExpense + "\n")

		if err != nil {
			return err
		}
	}

	// Successful write
	return nil
}

func WriteToConsole(formattedExpenses []string) {
	for _, formattedExpense := range formattedExpenses {
		fmt.Println(formattedExpense)
	}
}
