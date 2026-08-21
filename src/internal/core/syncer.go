package core

import (
	"bytes"
	"dsync/pkg/files"
	"dsync/pkg/goroutines"
	"dsync/pkg/logs"
	"dsync/pkg/sets"
	"fmt"
	"os"
	"path/filepath"
)

func SyncDirs(srcRoot string, srcDirs []string, dstRoot string, dstDirs []string) {
	context := goroutines.New()

	for _, newDir := range sets.SetDifference(srcDirs, dstDirs) {
		context.Add(func() {
			srcPath := filepath.Join(srcRoot, newDir)
			srcStat, err := os.Stat(srcPath)
			if err != nil {
				message := fmt.Sprintf("couldn't visit '%s'", srcPath)
				logs.Warning("core.SyncDirs", message, err)
				return
			}

			newPath := filepath.Join(dstRoot, newDir)
			err = os.MkdirAll(newPath, srcStat.Mode().Perm())
			if err != nil {
				message := fmt.Sprintf("couldn't create '%s'", newPath)
				logs.Warning("core.SyncDirs", message, err)
				return
			}

			fmt.Printf("+ '%s'\n", newDir)
		})
	}

	for _, oldDir := range sets.SetDifference(dstDirs, srcDirs) {
		context.Add(func() {
			oldPath := filepath.Join(dstRoot, oldDir)
			err := os.RemoveAll(oldPath)
			if err != nil {
				message := fmt.Sprintf("couldn't remove '%s'", oldPath)
				logs.Warning("core.SyncDirs", message, err)
				return
			}

			fmt.Printf("- '%s'\n", oldDir)
		})
	}

	context.Wait()
}

func SyncFiles(srcRoot string, srcFiles []string, dstRoot string, dstFiles []string) {
	context := goroutines.New()

	for _, newFile := range sets.SetDifference(srcFiles, dstFiles) {
		context.Add(func() {
			srcPath := filepath.Join(srcRoot, newFile)
			newPath := filepath.Join(dstRoot, newFile)
			err := files.Copy(srcPath, newPath)
			if err != nil {
				message := fmt.Sprintf("couldn't copy '%s'", newFile)
				logs.Warning("core.SyncFiles", message, err)
				return
			}

			fmt.Printf("+ '%s'\n", newFile)
		})
	}

	for _, oldFile := range sets.SetDifference(dstFiles, srcFiles) {
		context.Add(func() {
			oldPath := filepath.Join(dstRoot, oldFile)
			err := os.Remove(oldPath)
			if err != nil {
				message := fmt.Sprintf("couldn't remove '%s'", oldFile)
				logs.Warning("core.SyncFiles", message, err)
				return
			}

			fmt.Printf("- '%s'\n", oldFile)
		})
	}

	for _, existingFile := range sets.Intersection(srcFiles, dstFiles) {
		context.Add(func() {
			srcPath := filepath.Join(srcRoot, existingFile)
			srcChecksum, err := files.Checksum(srcPath)
			if err != nil {
				message := fmt.Sprintf("couldn't checksum '%s'", existingFile)
				logs.Warning("core.SyncFiles", message, err)
				return
			}

			dstPath := filepath.Join(dstRoot, existingFile)
			dstChecksum, err := files.Checksum(dstPath)
			if err != nil {
				message := fmt.Sprintf("couldn't checksum '%s'", existingFile)
				logs.Warning("core.SyncFiles", message, err)
				return
			}

			if bytes.Equal(srcChecksum, dstChecksum) {
				return
			}

			err = files.Copy(srcPath, dstPath)
			if err != nil {
				message := fmt.Sprintf("couldn't update '%s'", existingFile)
				logs.Warning("core.SyncFiles", message, err)
				return
			}

			fmt.Printf("M '%s'\n", existingFile)
		})
	}

	context.Wait()
}
