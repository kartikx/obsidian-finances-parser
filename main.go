package main

import (
	"fmt"
	"github.com/kartikx/obsidian-finances-parser/parser"
)

func main() {
	expenses, err := parser.ParseStatement(FinanceStatementFilePath, FinanceStatementFormat)

	if err != nil {
		fmt.Println("Parsing Statement failed", err)
		return
	}

	formattedExpenses, err := formatExpensesForObsidian(expenses)

	if err != nil {
		fmt.Println("Formatting expenses failed", err)
		return
	}

	if len(OutputFilePath) > 0 {
		writeToObsidianVault(formattedExpenses, OutputFilePath)
	} else {
		writeToConsole(formattedExpenses)
	}
}
