package core

import (
	"dsync/pkg/log"
	"fmt"
	"io/fs"
	"path/filepath"
)

func Scan(root string) ([]string, []string) {
	var dirs []string
	var files []string

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			log.Warning(fmt.Sprintf("unable to scan '%s'", path), err)
			return nil
		}

		if root == path {
			return nil
		}

		path, err = filepath.Rel(root, path)
		if err != nil {
			log.Warning(fmt.Sprintf("unable to process '%s'", path), err)
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
