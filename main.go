package main

import (
	"github.com/kartikx/obsidian-finances-parser/pkg/cmd"
)

func main() {
	// run obsidian-finances-parser
	obsidianParserCommand := cmd.NewObsidianParserCommand()
	obsidianParserCommand.Execute()
}
