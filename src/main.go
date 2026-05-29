package main

import (
	"dsync/cmd"
	"dsync/internal/core"
)

func main() {
	srcRoot := cmd.GetSrcRoot()
	dstRoot := cmd.GetDstRoot()

	srcDirs := core.ScanDirs(srcRoot)
	dstDirs := core.ScanDirs(dstRoot)
	core.SyncDirs(srcRoot, srcDirs, dstRoot, dstDirs)

	srcFiles := core.ScanFiles(srcRoot)
	dstFiles := core.ScanFiles(dstRoot)
	core.SyncFiles(srcRoot, srcFiles, dstRoot, dstFiles)
}
