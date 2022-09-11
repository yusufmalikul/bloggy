package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bloggy",
	Short: "Bloggy is a simple and stup*d blog generator",
	Long:  `A simple and stup*d blog generator written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bloggy")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
