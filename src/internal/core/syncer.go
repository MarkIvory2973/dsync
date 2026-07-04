package core

import (
	"bytes"
	"dsync/pkg/files"
	"dsync/pkg/log"
	"dsync/pkg/sets"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func SyncDirs(srcRoot string, srcDirs []string, dstRoot string, dstDirs []string) {
	var waitGroup sync.WaitGroup

	for _, newDir := range sets.SetDifference(srcDirs, dstDirs) {
		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, newDir)
			srcStat, err := os.Stat(srcPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to visit '%s'", srcPath), err)
				return
			}

			newPath := filepath.Join(dstRoot, newDir)
			err = os.MkdirAll(newPath, srcStat.Mode().Perm())
			if err != nil {
				log.Warning(fmt.Sprintf("unable to create '%s'", newPath), err)
				return
			}

			fmt.Printf("+ '%s'\n", newDir)
		})
	}

	for _, oldDir := range sets.SetDifference(dstDirs, srcDirs) {
		waitGroup.Go(func() {
			oldPath := filepath.Join(dstRoot, oldDir)
			err := os.RemoveAll(oldPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to remove '%s'", oldPath), err)
				return
			}

			fmt.Printf("- '%s'\n", oldDir)
		})
	}

	waitGroup.Wait()
}

func SyncFiles(srcRoot string, srcFiles []string, dstRoot string, dstFiles []string) {
	var waitGroup sync.WaitGroup

	for _, newFile := range sets.SetDifference(srcFiles, dstFiles) {
		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, newFile)
			newPath := filepath.Join(dstRoot, newFile)
			err := files.Copy(srcPath, newPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to copy '%s'", newFile), err)
				return
			}

			fmt.Printf("+ '%s'\n", newFile)
		})
	}

	for _, oldFile := range sets.SetDifference(dstFiles, srcFiles) {
		waitGroup.Go(func() {
			oldPath := filepath.Join(dstRoot, oldFile)
			err := os.Remove(oldPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to remove '%s'", oldFile), err)
				return
			}

			fmt.Printf("- '%s'\n", oldFile)
		})
	}

	for _, existingFile := range sets.Intersection(srcFiles, dstFiles) {
		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, existingFile)
			srcChecksum, err := files.Checksum(srcPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to checksum '%s'", existingFile), err)
				return
			}

			dstPath := filepath.Join(dstRoot, existingFile)
			dstChecksum, err := files.Checksum(dstPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to checksum '%s'", existingFile), err)
				return
			}

			if bytes.Equal(srcChecksum, dstChecksum) {
				return
			}

			err = files.Copy(srcPath, dstPath)
			if err != nil {
				log.Warning(fmt.Sprintf("unable to update '%s'", existingFile), err)
				return
			}

			fmt.Printf("M '%s'\n", existingFile)
		})
	}

	waitGroup.Wait()
}
