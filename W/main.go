package main

import "fmt"
import "os"
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "W",
	Short: "a javascript syntax tree walker",
	Long:  `a Takes a look at javascript code and tells us something ... `,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			return
		}
		rootDir := args[0]

		if err := parsePrograms(rootDir); err != nil {
			panic(err)
		}

	},
}

func main() {
	fmt.Println("testing....")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
