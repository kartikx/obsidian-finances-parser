/**
 * Contains methods to generate and write data for Obsidian.
 */

package main

func formatExpensesForObsidian(expenses []expense) ([]string, error) {
	// Converts each expense type into a formatted expense.
	// Look at README for the format.
	// Keep "categories" as empty for now. Eventually build logic to parse it out of name.

	return nil, nil
}

func writeToObsidianVault(formattedExpenses []string, outputFilePath string) {
	// Write to the path.
}

func writeToConsole(formattedExpenses []string) {
	// Just write to console one-by-one.
}
