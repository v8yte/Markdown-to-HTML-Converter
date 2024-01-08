/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

    "github.com/gomarkdown/markdown"
//	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/spf13/cobra"
)

// markdownCmd represents the markdown command
var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "file-converter markdown <inputfile> <outputfile>",
    Args: cobra.MatchAll(cobra.MinimumNArgs(1),cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
        inputfile := args[0]
        outputfile := "output.html"
        if len(args) == 2 {
            outputfile = args[1]
        }

        md, err := os.ReadFile(inputfile)
        if err != nil {
            fmt.Printf("Error reading file: %s\n", err)
            os.Exit(1)
        }

        html := mdToHTML(md)

        err = os.WriteFile(outputfile, html, 0644)
        if err != nil {
            fmt.Printf("Error reading file: %s\n", err)
            os.Exit(1)
        }

	},
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func init() {
	rootCmd.AddCommand(markdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// markdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// markdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
