package main

import "fmt"
import "os"
import "github.com/spf13/cobra"
import "path/filepath"

func isJS(path string) bool {
	return filepath.Ext(path) == ".js"
}

func listDir(root string) (paths []string, err error) {

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			abs, err := filepath.Abs(path)
			if err != nil {
				return err
			}

			if isJS(abs) {
				paths = append(paths, abs)
			}
		}
		return nil
	})

	return
}

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

		paths, err := listDir(rootDir)
		if err != nil {
			panic(err)
		}
		for _, v := range paths {
			fmt.Println(v)
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
