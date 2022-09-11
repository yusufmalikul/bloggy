package cmd

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
)

var rootCmd = &cobra.Command{
	Use:   "bloggy",
	Short: "Bloggy is a simple and stup*d blog generator",
	Long:  `A simple and stup*d blog generator written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generate() {
	fmt.Println("Generating...")
	sourceDir := "examples/posts"
	destinationDir := "examples/html"

	// iterate over the directory
	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		// read the file
		data, err := os.ReadFile(sourceDir + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		// generate the html
		var buf bytes.Buffer
		if err = goldmark.Convert(data, &buf); err != nil {
			fmt.Println(err)
		}

		// write the html
		err = os.WriteFile(destinationDir+"/"+file.Name()+".html", buf.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Done!")
}
