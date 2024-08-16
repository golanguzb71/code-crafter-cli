package commands

import (
	"codeCrafter/cmd"
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLi Tool",
	Run:   version,
}

func version(cmd *cobra.Command, args []string) {
	fmt.Println("Cli Tool version : 1.0.0")
}

func init() {
	cmd.RootCmd.AddCommand(versionCmd)
}
