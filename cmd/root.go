package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/yusufmalikul/bloggy/pkg/slug"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/spf13/cobra"
	"github.com/yuin/goldmark"
)

func init() {
	rootCmd.Flags().String("content", "", "Content directory")
	rootCmd.Flags().String("layouts", "", "Layouts directory")
	rootCmd.Flags().String("output", "", "Output directory")
	rootCmd.MarkFlagRequired("content")
	rootCmd.MarkFlagRequired("layouts")
	rootCmd.MarkFlagRequired("output")
}

var rootCmd = &cobra.Command{
	Use:   "bloggy",
	Short: "Bloggy is a simple and stup*d blog generator",
	Long:  `A simple and stup*d static blog generator written in Go (https://github.com/yusufmalikul/bloggy).`,
	Run: func(cmd *cobra.Command, args []string) {

		// get flags
		content, _ := cmd.Flags().GetString("content")
		layouts, _ := cmd.Flags().GetString("layouts")
		output, _ := cmd.Flags().GetString("output")

		generate(content, layouts, output)
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

func generate(content, layouts, output string) {
	fmt.Println("Generating...")
	t := time.Now()
	var posts Posts

	// iterate over the directory
	files, err := os.ReadDir(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	caser := cases.Title(language.English)
	for _, file := range files {
		// read the file
		data, err := os.ReadFile(content + "/" + file.Name())
		if err != nil {
			fmt.Println(err)
			return
		}

		// generate the html
		var buf bytes.Buffer
		if err = goldmark.Convert(data, &buf); err != nil {
			fmt.Println(err)
			return
		}

		// create output dir if not exist
		err = os.MkdirAll(output, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		// remove .md from the file name
		fileName := file.Name()
		fileName = fileName[:len(fileName)-3]

		// Store post title and body
		posts.Posts = append(posts.Posts, Post{Title: caser.String(strings.ReplaceAll(slug.Slugify(fileName), "-", " ")), Body: buf.String(), Path: slug.Slugify(fileName)})
	}

	// generate the index
	indexLayout := layouts + "/index.html"
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
	destinationIndex := output + "/index.html"
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

	// generate the post
	postLayout := layouts + "/post.html"
	postTemplate, err := os.ReadFile(postLayout)
	if err != nil {
		fmt.Println(err)
		return
	}

	// compile the template
	tmpl, err = template.New("post").Parse(string(postTemplate))
	if err != nil {
		fmt.Println(err)
		return
	}

	// write the template
	for _, post := range posts.Posts {

		// create the post dir
		err = os.MkdirAll(output+"/"+post.Path, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		destinationPost := path.Join(output, post.Path, "index.html")
		f, err := os.Create(destinationPost)
		if err != nil {
			fmt.Println(err)
			return
		}

		// execute the template
		err = tmpl.Execute(f, post)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Done in", time.Since(t))
}
