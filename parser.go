/*
 * Contains methods to parse statements from multiple banks into a unified format
 * that can be written into Obsidian.
 */

package main

type expense struct {
	name   string
	amount string
	date   string // Should this be a Date type instead?
}

func parseStatement(financeStatementFilePath, financeStatementFormat string) ([]expense, error) {
	// TODO Switch based on financeStatementFormat.

	var expenses []expense = nil
	return expenses, nil
}
