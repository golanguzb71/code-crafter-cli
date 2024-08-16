package main

import (
	"codeCrafter/cmd"
	_ "codeCrafter/cmd/commands"
)

func main() {
	cmd.Execute()
}
