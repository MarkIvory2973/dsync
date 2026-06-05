package core

import (
	"bytes"
	"dsync/pkg/files"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync"
)

func SyncDirs(srcRoot string, srcDirs []string, dstRoot string, dstDirs []string) {
	var waitGroup sync.WaitGroup

	for _, srcDir := range srcDirs {
		if slices.Contains(dstDirs, srcDir) {
			continue
		}

		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, srcDir)
			srcStat, err := os.Stat(srcPath)
			if err != nil {
				fmt.Printf("warning: unable to visit '%s': %v\n", srcPath, err)
				return
			}

			newPath := filepath.Join(dstRoot, srcDir)
			err = os.MkdirAll(newPath, srcStat.Mode().Perm())
			if err != nil {
				fmt.Printf("warning: unable to create '%s': %v\n", newPath, err)
				return
			}

			fmt.Printf("+ '%s'\n", srcDir)
		})
	}

	for _, dstDir := range dstDirs {
		if slices.Contains(srcDirs, dstDir) {
			continue
		}

		waitGroup.Go(func() {
			delPath := filepath.Join(dstRoot, dstDir)
			err := os.RemoveAll(delPath)
			if err != nil {
				fmt.Printf("warning: unable to remove '%s': %v\n", delPath, err)
				return
			}

			fmt.Printf("- '%s'\n", dstDir)
		})
	}

	waitGroup.Wait()
}

func SyncFiles(srcRoot string, srcFiles []string, dstRoot string, dstFiles []string) {
	var waitGroup sync.WaitGroup

	for _, srcFile := range srcFiles {
		if slices.Contains(dstFiles, srcFile) {
			continue
		}

		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, srcFile)
			dstPath := filepath.Join(dstRoot, srcFile)
			err := files.Copy(srcPath, dstPath)
			if err != nil {
				fmt.Printf("warning: unable to copy '%s': %v\n", srcFile, err)
				return
			}

			fmt.Printf("+ '%s'\n", srcFile)
		})
	}

	for _, dstFile := range dstFiles {
		if slices.Contains(srcFiles, dstFile) {
			continue
		}

		waitGroup.Go(func() {
			delPath := filepath.Join(dstRoot, dstFile)
			err := os.Remove(delPath)
			if err != nil {
				fmt.Printf("warning: unable to remove '%s': %v\n", dstFile, err)
				return
			}

			fmt.Printf("- '%s'\n", dstFile)
		})
	}

	for _, srcFile := range srcFiles {
		if !slices.Contains(dstFiles, srcFile) {
			continue
		}

		waitGroup.Go(func() {
			srcPath := filepath.Join(srcRoot, srcFile)
			srcChecksum, err := files.Checksum(srcPath)
			if err != nil {
				fmt.Printf("warning: unable to checksum '%s': %v\n", srcFile, err)
				return
			}

			dstPath := filepath.Join(dstRoot, srcFile)
			dstChecksum, err := files.Checksum(dstPath)
			if err != nil {
				fmt.Printf("warning: unable to checksum '%s': %v\n", srcFile, err)
				return
			}

			if bytes.Equal(srcChecksum, dstChecksum) {
				return
			}

			err = files.Copy(srcPath, dstPath)
			if err != nil {
				fmt.Printf("warning: unable to update '%s': %v\n", srcFile, err)
				return
			}

			fmt.Printf("M '%s'\n", srcFile)
		})
	}

	waitGroup.Wait()
}
