package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Act as a server",
	Long:  `You can run bloggy as a server.`,
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func server() {
	// serve static files on port 8080
	fmt.Println("Listening to port 8080...")
	fmt.Println("Visit http://localhost:8080")
	dir := "examples/html"
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
