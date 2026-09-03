//go:build !windows
// +build !windows

// Build Constraints : https://golang.org/pkg/go/build/#hdr-Build_Constraints
package main

import (
	"fmt"
	"os"
	"syscall" // using syscall is a bad idea it is deprecated in favor of "golang.org/x/sys/unix"
)

func main() {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error doing os.Getwd() err:", err)
		os.Exit(1)
	}

	err = syscall.Statfs(wd, &stat)
	if err != nil {
		fmt.Println("Error doing syscall.Statfs(wd, &stat) err:", err)
		os.Exit(1)
	}
	fmt.Printf("Current working directory is : %v\n", wd)
	fmt.Println("# using syscall.Statfs : http://man7.org/linux/man-pages/man2/statfs.2.html")
	fmt.Printf("Maximum lenght file name : %d \n", stat.Namelen)
	fmt.Printf("# Free blocks available to unprivileged * block size in Byte (%v) gives bytes available on disk\n", uint64(stat.Bsize))

	bytesTotal := stat.Blocks * uint64(stat.Bsize)
	bytesFree := stat.Bavail * uint64(stat.Bsize)
	bytesUsed := bytesTotal - bytesFree
	percentUsed := float32(bytesUsed) / float32(bytesTotal) * 100
	mbTotal := bytesTotal / 1024 / 1024
	mbFree := bytesFree / 1024 / 1024
	mbUsed := mbTotal - mbFree
	fmt.Printf("%v Blocks available \t(from Total of %v Blocks)\t used: %v Block \n",
		stat.Bavail, stat.Blocks, stat.Blocks-stat.Bavail)
	fmt.Printf("%v inodes available \t(from Total of %v inodes)\t used: %v inodes \n",
		stat.Ffree, stat.Files, stat.Files-stat.Ffree)
	fmt.Printf("%v Bytes available \t(from Total of %v Bytes)\t (%.2f)%% used: %v \n",
		bytesFree, bytesTotal, percentUsed, bytesUsed)
	fmt.Printf("%v KB available \t(from Total of %v KB)\t %.2f%% used: %v KB\n",
		bytesFree/1024, bytesTotal/1024, percentUsed, bytesUsed/1024)
	fmt.Printf("%v MB available \t(from Total of %v MB)\t %.2f%% used: %v MB\n",
		mbFree, mbTotal, percentUsed, mbUsed)
	fmt.Printf("%v GB available \t(from Total of %v GB)\t\t %.2f%% used: %v GB\n",
		mbFree/1024, mbTotal/1024, percentUsed, mbUsed/1024)

}
