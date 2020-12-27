package main

import "path/filepath"
import "github.com/peter9207/F/W/javascript"
import "os"
import "io/ioutil"

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
		file, err := ioutil.ReadFile(v)
		if err != nil {
			return err
		}

		err = javascript.Parse(string(file))
		if err != nil {
			return err
		}
	}
	return
}
