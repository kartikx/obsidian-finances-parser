// Contains methods to generate and write data for Obsidian.
package obsidian

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/kartikx/obsidian-finances-parser/pkg/models"
)

var dateAttributeInFormattedExpenseRegex = regexp.MustCompile(`\(date::(.*?)\)`)

func formatExpenseForObsidian(expense *models.Expense) string {
	categories := ""

	for _, category := range expense.Categories {
		categories += fmt.Sprintf("\"%s\",", category.String())
	}

	if len(categories) > 0 {
		categories = categories[:len(categories)-1]
	} else {
		categories = "\"" + models.UNKNOWN_EXPENSE.String() + "\""
	}

	return fmt.Sprintf("- #expense (name::%s) (amount::%.2f) (date::%s) (categories::%s) (note::%s)",
		expense.Name,
		expense.Amount,
		expense.Date,
		categories,
		expense.Note)
}

// TODO Rename this function and file.
func WriteToObsidianVault(expenses []*models.Expense, outputFilePath string) error {
	fmt.Println("Writing to Obsidian Vault")

	if len(outputFilePath) == 0 {
		return fmt.Errorf("output file path provided is empty")
	}

	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	fileInfo, err := os.Stat(outputFilePath)
	if err != nil {
		return err
	}

	// Every expense will definitely be in the future from this date.
	lastWrittenExpenseDate := "2023-01-01"

	// If expenses are already present in this file, find the last written expense.
	if fileSize := fileInfo.Size(); fileSize > 0 {
		lastWrittenExpense, err := findLastExpenseInFile(file, fileSize)

		if err != nil {
			return err
		}

		fmt.Println("Last already written expense is: ", lastWrittenExpense)

		lastWrittenExpenseDate, err = getDateFromFormattedExpense(lastWrittenExpense)

		if err != nil {
			lastWrittenExpenseDate = "2023-01-01"
			fmt.Printf("Unable to parse date from last line: \"%s\". Using default date \"%s\" instead",
				lastWrittenExpense, lastWrittenExpenseDate)
		}
	}

	for _, expense := range expenses {
		// Skip forward to find the first transaction that has not been written, based on the date.
		// TODO @kartikx This might skip some transactions on the last date, which differs only in the timestamp.
		if lastWrittenExpenseDate >= expense.Date {
			continue
		}

		formattedExpense := formatExpenseForObsidian(expense)

		_, err = file.WriteString(formattedExpense + "\n")

		if err != nil {
			return err
		}
	}

	// Successful write
	return nil
}

func WriteToConsole(expenses []*models.Expense) {
	for _, expense := range expenses {
		fmt.Println(formatExpenseForObsidian(expense))
	}
}

func findLastExpenseInFile(file *os.File, fileSize int64) (string, error) {
	lastExpense := ""

	// An expense line should be 300 bytes at maximum.
	// Seeking 512 bytes from the end should provide at least one expense.
	bytesToSeek := math.Min(512, float64(fileSize))

	_, err := file.Seek(int64(-bytesToSeek), io.SeekEnd)

	if err != nil {
		return "", err
	}

	buffer := make([]byte, 512)

	bytesRead, err := file.Read(buffer)

	if err != nil {
		return "", err
	}

	lines := strings.Split(string(buffer[:bytesRead]), "\n")

	// Skip trailing new lines.
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])

		if len(line) > 0 {
			lastExpense = line
			break
		}
	}

	return lastExpense, nil
}

func getDateFromFormattedExpense(formattedExpense string) (string, error) {
	matches := dateAttributeInFormattedExpenseRegex.FindStringSubmatch(formattedExpense)

	if matches == nil || len(matches) < 2 {
		return "", fmt.Errorf("no date found in expense: %s", formattedExpense)
	}

	return matches[1], nil
}
