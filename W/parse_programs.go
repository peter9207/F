package main

import "path/filepath"
import "github.com/peter9207/F/W/parser"
import "os"

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

func parsePrograms(root string) (err error) {

	paths, err := listDir(root)
	if err != nil {
		panic(err)
	}

	for _, v := range paths {
		err := parser.Parse(v)
		if err != nil {
			return err
		}

	}
	return
}
