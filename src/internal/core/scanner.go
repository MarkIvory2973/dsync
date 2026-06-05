package core

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Scan(root string) ([]string, []string) {
	var dirs []string
	var files []string

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("warning: unable to scan '%s': %v\n", path, err)
			return nil
		}

		if root == path {
			return nil
		}

		path, err = filepath.Rel(root, path)
		if err != nil {
			fmt.Printf("warning: unable to process '%s': %v\n", path, err)
			return nil
		}

		if entry.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}

		return nil
	})

	return dirs, files
}
