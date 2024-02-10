package main

import (
	"fmt"
)

func main() {
	// Read these as a CLI argument.
	financeStatementFilePath := "finances.pdf"
	outputFilePath := "Budget Planning/2024-01.md"

	// This should be an Enum.
	financeStatementFormat := "HDFC-DEBIT"

	expenses, err := parseStatement(financeStatementFilePath, financeStatementFormat)

	if err != nil {
		fmt.Println("Parsing Statement failed", err)
		return
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
