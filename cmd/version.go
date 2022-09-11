package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version        = ""
	Commit         = ""
	BuildTimestamp = ""
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bloggy",
	Long: `You can see the version of bloggy.
			This is usually needed when you want to report a bug.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
		fmt.Println(Commit)
		fmt.Println(BuildTimestamp)
	},
}
