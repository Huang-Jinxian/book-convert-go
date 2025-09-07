package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	title           string
	author          string
	coverPath       string
	chapterTitleReg string
	fileEncoding    string
	filePath        string
	destPath        string

	port int16
)

var rootCmd = &cobra.Command{
	Use:   "book_convert",
	Short: "Convert book between txt，epub and azw3",
	Long: `Book Convert is a CLI that helps you to convert your book.
Support book type includes txt, epub and azw3.`,
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert txt book to epub book",
	Long:  `Convert txt book to epub book.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   convert,
}

var webCmd = &cobra.Command{
	Use:   "web [-p port]",
	Short: "start web server",
	Long:  `#todo need to implement the webServer func`,
	Args:  cobra.MinimumNArgs(1),
	Run:   webServer,
}

func init() {

	convertCmd.Flags().StringVarP(&title, "title", "t", "", "book title")
	convertCmd.Flags().StringVarP(&author, "author", "a", "", "book author")
	convertCmd.Flags().StringVarP(&coverPath, "cover", "p", "", "cover path")
	convertCmd.Flags().StringVarP(&chapterTitleReg, "chapter", "r", "", "chapter title regular expression")
	convertCmd.Flags().StringVarP(&fileEncoding, "encoding", "e", "GB18030", "encoding type")
	convertCmd.Flags().StringVarP(&filePath, "file", "f", "", "file path")
	convertCmd.Flags().StringVarP(&destPath, "output", "o", "", "output directory")

	webCmd.Flags().Int16VarP(&port, "port", "p", 8080, "Port to listen on")

	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(webCmd)
}

func convert(cmd *cobra.Command, args []string) {
	fmt.Println(title, author, coverPath, chapterTitleReg, fileEncoding, filePath, destPath)
}

func webServer(cmd *cobra.Command, args []string) {

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
