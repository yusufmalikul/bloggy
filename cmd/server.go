package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().String("dir", "", "Directory to serve")
	serverCmd.MarkFlagRequired("dir")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Act as a server",
	Long:  `You can run bloggy as a server.`,
	Run: func(cmd *cobra.Command, args []string) {
		// get flags
		dir, _ := cmd.Flags().GetString("dir")
		server(dir)
	},
}

func server(dir string) {
	// serve static files on port 8080
	fmt.Println("Listening...")
	fmt.Println("Visit http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
