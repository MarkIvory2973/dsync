package core

import (
	"dsync/pkg/logs"
	"fmt"
	"io/fs"
	"path/filepath"
)

func Scan(root string) ([]string, []string) {
	var dirs []string
	var files []string

	filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			message := fmt.Sprintf("couldn't scan '%s'", path)
			logs.Warning("core.Scan", message, err)
			return nil
		}

		if root == path {
			return nil
		}

		path, err = filepath.Rel(root, path)
		if err != nil {
			message := fmt.Sprintf("couldn't process '%s'", path)
			logs.Warning("core.Scan", message, err)
			return nil
		}

		if entry.IsDir() {
			dirs = append(dirs, path)
		} else if entry.Type().IsRegular() {
			files = append(files, path)
		}

		return nil
	})

	return dirs, files
}
