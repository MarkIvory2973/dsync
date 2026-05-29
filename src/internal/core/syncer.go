package core

import (
	"bytes"
	"dsync/pkg/files"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func SyncDirs(srcRoot string, srcDirs []string, dstRoot string, dstDirs []string) {
	for _, srcDir := range srcDirs {
		if slices.Contains(dstDirs, srcDir) {
			continue
		}

		srcPath := filepath.Join(srcRoot, srcDir)
		srcStat, err := os.Stat(srcPath)
		if err != nil {
			fmt.Printf("warning: unable to visit '%s': %v\n", srcPath, err)
			continue
		}

		newPath := filepath.Join(dstRoot, srcDir)
		err = os.MkdirAll(newPath, srcStat.Mode().Perm())
		if err != nil {
			fmt.Printf("warning: unable to create '%s': %v\n", newPath, err)
			continue
		}

		fmt.Printf("+ '%s'\n", srcDir)
	}

	for _, dstDir := range dstDirs {
		if slices.Contains(srcDirs, dstDir) {
			continue
		}

		delPath := filepath.Join(dstRoot, dstDir)
		err := os.RemoveAll(delPath)
		if err != nil {
			fmt.Printf("warning: unable to remove '%s': %v\n", delPath, err)
			continue
		}

		fmt.Printf("- '%s'\n", dstDir)
	}
}

func SyncFiles(srcRoot string, srcFiles []string, dstRoot string, dstFiles []string) {
	for _, srcFile := range srcFiles {
		if slices.Contains(dstFiles, srcFile) {
			continue
		}

		srcPath := filepath.Join(srcRoot, srcFile)
		dstPath := filepath.Join(dstRoot, srcFile)
		err := files.Copy(srcPath, dstPath)
		if err != nil {
			fmt.Printf("warning: unable to copy '%s': %v\n", srcFile, err)
			continue
		}

		fmt.Printf("+ '%s'\n", srcFile)
	}

	for _, dstFile := range dstFiles {
		if slices.Contains(srcFiles, dstFile) {
			continue
		}

		delPath := filepath.Join(dstRoot, dstFile)
		err := os.Remove(delPath)
		if err != nil {
			fmt.Printf("warning: unable to remove '%s': %v\n", dstFile, err)
			continue
		}

		fmt.Printf("- '%s'\n", dstFile)
	}

	for _, srcFile := range srcFiles {
		if !slices.Contains(dstFiles, srcFile) {
			continue
		}

		srcPath := filepath.Join(srcRoot, srcFile)
		srcChecksum, err := files.Checksum(srcPath)
		if err != nil {
			fmt.Printf("warning: unable to checksum '%s': %v\n", srcFile, err)
			continue
		}

		dstPath := filepath.Join(dstRoot, srcFile)
		dstChecksum, err := files.Checksum(dstPath)
		if err != nil {
			fmt.Printf("warning: unable to checksum '%s': %v\n", srcFile, err)
			continue
		}

		if bytes.Equal(srcChecksum, dstChecksum) {
			continue
		}

		err = files.Copy(srcPath, dstPath)
		if err != nil {
			fmt.Printf("warning: unable to update '%s': %v\n", srcFile, err)
			continue
		}

		fmt.Printf("M '%s'\n", srcFile)
	}
}
