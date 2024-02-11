package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kartikx/obsidian-finances-parser/models"
)

func ParseHdfcStatement(statementPath string) ([]*models.Expense, error) {
	statement, err := readStatement(statementPath)

	if err != nil {
		return nil, err
	}

	// Skip first 2 lines
	expenseLines := strings.Split(statement, "\n")[2:]

	var expenses []*models.Expense

	for _, expenseLine := range expenseLines {
		expense, err := parseExpense(expenseLine)

		if err != nil {
			err = fmt.Errorf("parsing expenses on page failed %s", err.Error())
			return expenses, err
		}

		if expense != nil {
			expenses = append(expenses, expense)
		}
	}

	return expenses, nil
}

func readStatement(statementPath string) (string, error) {
	content, err := os.ReadFile(statementPath)

	if err != nil {
		err = fmt.Errorf("reading finances.txt failed with %s", err.Error())
		return "", err
	}

	return string(content), nil
}

func parseExpense(expenseLine string) (*models.Expense, error) {
	var exp *models.Expense

	expenseLine = strings.TrimSpace(expenseLine)

	if len(expenseLine) == 0 {
		return nil, nil
	}

	// TODO Store a list of expenses that were skipped over, to allow parsing to complete but still track where it failed.
	// We wouldn't want to throw an error on the first failure.

	fields, err := parseCsvExpenseLine(expenseLine)

	if err != nil {
		return nil, err
	}

	exp, err = constructExpenseFromCsvFields(fields)

	categories := GetCategoriesForExpense(exp)
	exp.Categories = categories

	exp.Name, exp.Note = formatExpenseName(exp)

	// TODO Add Rules to filter the expense.

	return exp, nil
}

func constructExpenseFromCsvFields(fields []string) (*models.Expense, error) {
	date := convertSlashDateToIso(fields[0])

	name := strings.TrimSpace(fields[1])

	amount, err := getAmountFromDebitAndCreditAmount(fields[3], fields[4])

	if err != nil {
		return nil, err
	}

	return &models.Expense{
		Name:       name,
		Amount:     amount,
		Date:       date,
		Categories: []models.ExpenseCategory{},
	}, nil
}

func parseCsvExpenseLine(expenseLine string) ([]string, error) {
	fields := strings.Split(expenseLine, ",")

	if len(fields) != 7 {
		return nil, fmt.Errorf("unable to parse expenseLine %s into CSV", expenseLine)
	}

	return fields, nil
}

/*
 * Converts a slashed date (such as 07/01/24) into ISO format (2024-01-07)
 */
func convertSlashDateToIso(slashDate string) string {
	slashDate = strings.TrimSpace(slashDate)

	return "20" + slashDate[6:8] + "-" + slashDate[3:5] + "-" + slashDate[0:2]
}

/*
 * Remove unecessary fields from Expense such as BANK NAME.
 */
func formatExpenseName(expense *models.Expense) (string, string) {
	expenseName := strings.TrimSpace(expense.Name)
	expenseNote := ""

	if expenseName[0:3] == "UPI" {
		expenseNameFields := strings.Split(expenseName, "-")

		if expense.Amount > 0 {
			expenseName = fmt.Sprintf("TO (%s) FOR (%s)",
				expenseNameFields[1],
				expenseNameFields[len(expenseNameFields)-1])
		} else {
			expenseName = fmt.Sprintf("FROM (%s) FOR (%s)",
				expenseNameFields[1],
				expenseNameFields[len(expenseNameFields)-1])
		}
		expenseNote = fmt.Sprintf("UPI %s", expenseNameFields[2])
	} else if expenseName[0:3] == "ATW" {
		expenseNameFields := strings.Split(expenseName, "-")

		expenseName = fmt.Sprintf("ATM at (%s)",
			expenseNameFields[len(expenseNameFields)-1])
	} else if expenseName[0:4] == "NEFT" {
		expenseNameFields := strings.Split(expenseName, "-")
		expenseName = fmt.Sprintf("FROM (%s) TO (%s)",
			expenseNameFields[2],
			expenseNameFields[3])
		expenseNote = "NEFT"
	}

	return expenseName, expenseNote
}

/*
 * If amount is debitted, keep it as it is.
 * If amount is creditted, set it as a negative value.
 * Returns the credit or debit value as a float32.
 */
func getAmountFromDebitAndCreditAmount(debitStr string, creditStr string) (float32, error) {
	debitStr = strings.TrimSpace(debitStr)
	creditStr = strings.TrimSpace(creditStr)

	debitAmount, debitErr := strconv.ParseFloat(debitStr, 32)
	creditAmount, creditErr := strconv.ParseFloat(creditStr, 32)

	if debitErr != nil && creditErr != nil {
		return 0, fmt.Errorf("parsing Debit Amount {%s} and Credit amount {%s} failed. Debit: {%s} Credit: {%s}",
			debitStr, creditStr, debitErr.Error(), creditErr.Error())
	}

	if debitErr != nil {
		return -1 * float32(creditAmount), nil
	}

	if creditErr != nil {
		return float32(debitAmount), nil
	}

	if creditAmount != 0 {
		return -1 * float32(creditAmount), nil
	}

	return float32(debitAmount), nil
}
