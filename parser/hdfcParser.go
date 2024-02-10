package parser

import (
	"errors"
	"fmt"
	"github.com/kartikx/obsidian-finances-parser/models"
	"regexp"
	"strings"
)

// Regular expression pattern to match the date on the start of each expense.
// Declared globally to prevent re-compilation.
var dateRegex = regexp.MustCompile(`^\s*(\d{2}/\d{2}/\d{2})`)

var lastExpenseRegex = regexp.MustCompile(`^\s*HDFC BANK LIMITED`)

var lastExpenseOnLastPageRegex = regexp.MustCompile(`^\s*STATEMENT SUMMARY :-`)

var lineSplitPattern = regexp.MustCompile(`\s{5,}`)

func ParseHdfcStatement(statement string) ([]models.Expense, error) {
	pages := strings.Split(statement, "\f")

	var expenses []models.Expense

	for index, page := range pages {
		lines := strings.Split(page, "\n")
		expensesOnPage, err := parseExpensesOnPage(lines, index == len(pages)-2)

		if err != nil {
			err = fmt.Errorf("parsing expenses on page failed %s", err.Error())
			return expenses, err
		}

		expenses = append(expenses, expensesOnPage...)
	}

	return expenses, nil
}

// TODO Handle expenses that cross over pages.
func parseExpensesOnPage(lines []string, isLastPage bool) ([]models.Expense, error) {
	currExpenseStartIndex := -1

	expenses := make([]models.Expense, 0)

	// Find the first expense
	for index, line := range lines {
		if match := dateRegex.FindStringSubmatch(line); match != nil {
			currExpenseStartIndex = index
			break
		}
	}

	if currExpenseStartIndex < 0 {
		// This page contains no expenses
		return expenses, nil
	}

	currExpenseEndIndex := currExpenseStartIndex + 1

	// TODO @kartikx Ensure that this handles end lines and EOF well.
	for currExpenseEndIndex <= len(lines) {
		currLine := lines[currExpenseEndIndex]

		if match := dateRegex.FindStringSubmatch(currLine); match != nil {
			// We found a new expense.
			expense, err := parseExpense(lines[currExpenseStartIndex:currExpenseEndIndex])

			if err != nil {
				err = fmt.Errorf("parsing expense failed: %s", err.Error())
				return expenses, err
			}

			expenses = append(expenses, expense)

			currExpenseStartIndex = currExpenseEndIndex
		} else {
			var match []string
			if isLastPage {
				match = lastExpenseOnLastPageRegex.FindStringSubmatch(currLine)
			} else {
				match = lastExpenseRegex.FindStringSubmatch(currLine)
			}

			if match != nil {
				// We have reached the last line.
				expense, err := parseExpense(lines[currExpenseStartIndex:currExpenseEndIndex])

				if err != nil {
					err = fmt.Errorf("parsing expense failed: %s", err.Error())
					return expenses, err
				}

				expenses = append(expenses, expense)
				break
			}
		}
		currExpenseEndIndex++
	}

	return expenses, nil
}

func parseExpense(lines []string) (models.Expense, error) {
	var exp models.Expense

	if len(lines) == 0 {
		return exp, errors.New("no lines for expense")
	}

	trimmedLine := strings.TrimSpace(lines[0])

	fields := lineSplitPattern.Split(trimmedLine, -1)

	exp = models.Expense{
		Name:       strings.TrimSpace(fields[1]),
		Amount:     fields[4],
		Date:       fields[0],
		Categories: make([]string, 0),
	}

	for index := 1; index < len(lines); index++ {
		line := strings.TrimSpace(lines[index])

		if len(line) == 0 {
			continue
		}

		exp.Name += line
	}

	exp.Categories = guessCategoriesFromExpenseName(exp.Name)

	return exp, nil
}
