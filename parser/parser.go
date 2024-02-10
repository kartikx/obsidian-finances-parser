/*
 * Contains methods to parse statements from multiple banks into a unified format
 * that can be written into Obsidian.
 */

package parser

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kartikx/obsidian-finances-parser/models"
	"os"
	"os/exec"
)

var debug bool = true

func convertPdfToTxt(filePath string) (string, error) {
	args := []string{
		"-layout", // Maintain (as best as possible) the original physical layout of the text.
		filePath,  // The input file.
		"-",       // Send the output to stdout.
	}
	cmd := exec.CommandContext(context.Background(), "pdftotext", args...)

	var buf bytes.Buffer
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func guessCategoriesFromExpenseName(expenseName string) []string {
	return make([]string, 0)
}

// For local testing only, avoids re-exec of `pdftotext` again and again
func readConvertedTxtFile() (string, error) {
	content, err := os.ReadFile("finances.txt")

	if err != nil {
		err = fmt.Errorf("reading finances.txt failed with %s", err.Error())
		return "", err
	}

	return string(content), nil
}

func ParseStatement(financeStatementFilePath string, financeStatementFormat models.StatementFormat) ([]models.Expense, error) {
	var financeStatementText string
	var err error

	if debug {
		// Skip Pdf to Txt conversion in debug mode.
		financeStatementText, err = readConvertedTxtFile()

		if err != nil {
			return nil, err
		}
	} else {
		financeStatementText, err = convertPdfToTxt(financeStatementFilePath)

		if err != nil {
			err = fmt.Errorf("PDF to TXT conversion failed with %s", err.Error())
			return nil, err
		}
	}

	var expenses []models.Expense

	switch financeStatementFormat {
	case models.HDFC_DEBIT:
		expenses, err = ParseHdfcStatement(financeStatementText)
	default:
		err = fmt.Errorf("financeStatemtFormat %d is invalid.", financeStatementFormat)
	}

	if err != nil {
		err = fmt.Errorf("parsing Statement failed with %s", err.Error())
		return nil, err
	}

	return expenses, nil
}
