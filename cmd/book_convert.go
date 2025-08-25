package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	port int16
)

var rootCmd = &cobra.Command{
	Use:   "book_convert",
	Short: "Convert book between txt，epub and azw3",
	Long: `Book Convert is a CLI that helps you to convert your book.
Support book type includes txt, epub and azw3.`,
}

var webCmd = &cobra.Command{
	Use:   "web [-p port]",
	Short: "start web server",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run:   webServer,
}

func init() {

	webCmd.Flags().Int16VarP(&port, "port", "p", 8080, "Port to listen on")
	rootCmd.AddCommand(webCmd)
}

func webServer(cmd *cobra.Command, args []string) {

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
