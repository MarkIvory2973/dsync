package main

import (
	"dsync/cmd"
	"dsync/internal/core"
)

func main() {
	srcRoot := cmd.GetSrcRoot()
	dstRoot := cmd.GetDstRoot()

	srcDirs, srcFiles := core.Scan(srcRoot)
	dstDirs, dstFiles := core.Scan(dstRoot)

	core.SyncDirs(srcRoot, srcDirs, dstRoot, dstDirs)
	core.SyncFiles(srcRoot, srcFiles, dstRoot, dstFiles)
}
