package main

import (
	"fmt"
	"os"

	"github.com/kartikx/obsidian-finances-parser/models"
	"github.com/kartikx/obsidian-finances-parser/parser"
)

func main() {
	// Read these as a CLI argument.
	financeStatementFilePath := "finances.pdf"
	outputFilePath := "Budget Planning/2024-01.md"
	financeStatementFormat := models.HDFC_DEBIT

	expenses, err := parser.ParseStatement(financeStatementFilePath, financeStatementFormat)

	if err != nil {
		fmt.Println("Parsing Statement failed", err)
		return
	}

	outputFile, _ := os.Create("expenses.txt")

	for _, expense := range expenses {
		outputFile.WriteString(expense.String() + "\n")
	}

	formattedExpenses, err := formatExpensesForObsidian(expenses)

	if err != nil {
		fmt.Println("Formatting expenses failed", err)
		return
	}

	if outputFilePath != "" {
		writeToObsidianVault(formattedExpenses, outputFilePath)
	} else {
		writeToConsole(formattedExpenses)
	}
}
