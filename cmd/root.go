package cmd

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/yusufmalikul/bloggy/pkg/slug"

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

type Post struct {
	Title string
	Body  string
	Path  string
}

type Posts struct {
	Posts []Post
}

func generate() {
	fmt.Println("Generating...")
	t := time.Now()
	var posts Posts
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
		// remove .md from the file name
		fileName := file.Name()
		fileName = fileName[:len(fileName)-3]
		err = os.WriteFile(destinationDir+"/"+fileName+".html", buf.Bytes(), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Store post title and body
		posts.Posts = append(posts.Posts, Post{Title: slug.Slugify(fileName), Body: buf.String(), Path: slug.Slugify(fileName) + ".html"})
	}

	// generate the index
	indexLayout := "examples/layouts/index.html"
	indexTemplate, err := os.ReadFile(indexLayout)
	if err != nil {
		fmt.Println(err)
		return
	}

	// compile the template
	tmpl, err := template.New("index").Parse(string(indexTemplate))
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the template
	destinationIndex := destinationDir + "/index.html"
	f, err := os.Create(destinationIndex)
	if err != nil {
		fmt.Println(err)
		return
	}

	// execute the template
	err = tmpl.Execute(f, posts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done in", time.Since(t))
}
