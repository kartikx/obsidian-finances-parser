package cmd

import (
	"fmt"
	"github.com/kartikx/obsidian-finances-parser/pkg/obsidian"
	"github.com/kartikx/obsidian-finances-parser/pkg/parser"
	"github.com/spf13/cobra"
)

func NewObsidianParserCommand() *cobra.Command {
	var inputFilePath, outputFilePath string
	var rootCmd = &cobra.Command{
		Use:   "ofp", // TODO: @renormalize @kartikx command name yet to be decided on
		Short: "Parse your bank statements",
		Long:  `This program attempts to convert HDFC bank statements into a format that my Obsidian budget planning workflow can understand.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO @renormalize: FinanceStatementFormat must be inferred from input file
			expenses, err := parser.ParseStatement(inputFilePath, FinanceStatementFormat)
			if err != nil {
				fmt.Println("Parsing Statement failed", err)
				return
			}

			formattedExpenses, err := obsidian.FormatExpensesForObsidian(expenses)

			if err != nil {
				fmt.Println("Formatting expenses failed", err)
				return
			}

			// TODO:could be improved to decide code flow based on state of the program instead of file path
			if len(outputFilePath) > 0 {
				obsidian.WriteToObsidianVault(formattedExpenses, outputFilePath)
			} else {
				obsidian.WriteToConsole(formattedExpenses)
			}
		},
	}
	// TODO @kartikx is the input file is always required from the user?
	// or is there a sane default that would be hardcoded in?
	rootCmd.PersistentFlags().StringVar(&inputFilePath, "input-file", FinanceStatementFilePath, "the file path of the finance statement")
	rootCmd.PersistentFlags().StringVar(&outputFilePath, "output-file", OutputFilePath, "the output file path, decides output to console or vault")
	// rootCmd.MarkFlagRequired("input-file")
	return rootCmd
}
