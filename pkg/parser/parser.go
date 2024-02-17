// Contains methods to parse statements from multiple banks into a unified format
// that can be written into Obsidian.

package parser

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kartikx/obsidian-finances-parser/pkg/models"
	// "os"
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

func ParseStatement(financeStatementFilePath string, financeStatementFormat models.StatementFormat) ([]*models.Expense, error) {
	var err error

	var expenses []*models.Expense

	switch financeStatementFormat {
	case models.HDFC_DEBIT:
		expenses, err = ParseHdfcStatement(financeStatementFilePath)
	default:
		err = fmt.Errorf("financeStatemtFormat %d is invalid.", financeStatementFormat)
	}

	if err != nil {
		err = fmt.Errorf("parsing Statement failed with %s", err.Error())
		return nil, err
	}

	return expenses, nil
}
