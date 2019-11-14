// +build !windows

// Build Constraints : https://golang.org/pkg/go/build/#hdr-Build_Constraints
package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"math"
	"os"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskUsage struct {
	TotalBytes  uint64  `json:"total"`
	UsedBytes   uint64  `json:"used"`
	AvailBytes  uint64  `json:"avail"`
	PercentUsed float32 `json:"percentUsed"`
	TotalInodes uint64  `json:"nodesTotal"`
	UsedInodes  uint64  `json:"nodesUsed"`
	AvailInodes uint64  `json:"nodesAvail"`
}

func GetDiskUsage(path string) (DiskUsage, error) {
	var stat unix.Statfs_t
	var du = DiskUsage{}

	err := unix.Statfs(path, &stat)
	if err != nil {
		return du, fmt.Errorf("GetDiskUsage: Error doing syscall.Statfs(wd, &stat) err: %v", err)
	}
	du.TotalBytes = stat.Blocks * uint64(stat.Bsize)
	du.AvailBytes = stat.Bavail * uint64(stat.Bsize)
	du.UsedBytes = du.TotalBytes - du.AvailBytes
	du.PercentUsed = float32(du.UsedBytes) / float32(du.TotalBytes) * 100
	du.TotalInodes = stat.Files
	du.AvailInodes = stat.Ffree
	du.UsedInodes = du.TotalInodes - du.AvailInodes
	return du, nil
}

func PrintMsgDiskUsage(du DiskUsage, sizeUnit uint64) {
	var template string
	switch sizeUnit {
	case B:
		template = "%.0f Bytes available \t(from Total of %.0f Bytes)\t %.2f%% used: %.0f Bytes\n"
	case KB:
		template = "%.0f KB available \t(from Total of %.0f KB)\t %.2f%% used: %.0f KB\n"
	case MB:
		template = "%.0f MB available \t(from Total of %.0f MB)\t %.2f%% used: %.0f MB\n"
	case GB:
		template = "%.0f GB available \t(from Total of %.0f GB)\t %.2f%% used: %.0f GB\n"
	default:
		panic(fmt.Sprintf("# ERROR IN PrintMsgDiskUsage INVALID sizeUnit: %v", sizeUnit))
	}
	fmt.Printf(template,
		math.Round(float64(du.AvailBytes)/float64(sizeUnit)),
		math.Round(float64(du.TotalBytes)/float64(sizeUnit)),
		du.PercentUsed,
		math.Round(float64(du.UsedBytes)/float64(sizeUnit)))
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("FATAL ERROR doing os.Getwd() err: %v", err))
	}

	du, err := GetDiskUsage(wd)
	if err != nil {
		panic(fmt.Sprintf("FATAL ERROR : %v", err))
	}

	fmt.Printf("Current working directory is : %v\n", wd)
	fmt.Println("# using unix.Statfs : http://man7.org/linux/man-pages/man2/statfs.2.html")

	fmt.Printf("%v inodes available \t(from Total of %v inodes)\t used: %v inodes \n",
		du.AvailInodes, du.TotalInodes, du.UsedInodes)

	PrintMsgDiskUsage(du, B)
	PrintMsgDiskUsage(du, KB)
	PrintMsgDiskUsage(du, MB)
	PrintMsgDiskUsage(du, GB)

}
