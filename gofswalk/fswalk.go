package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func stats(path string, fi os.FileInfo, err error) error {
	if err != nil { return nil }
	stat := fi.Sys().(*syscall.Stat_t)
	mtime := time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec))
	atime := time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec))
	ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	// path mode uid gid size atime mtime ctime
	fmt.Printf("%s,%o,%v,%v,%v,%v,%v,%v\n", path, fi.Mode(), stat.Uid, stat.Gid, fi.Size(), atime.Unix(), mtime.Unix(), ctime.Unix())
	return nil
}

func main() {
	// read arguments 
	// walk https://golang.org/pkg/path/filepath/#Walk
	// stat https://golang.org/src/os/stat.go?s=642:683#L10
	args := os.Args[1:]
	fmt.Println("path,mode,uid,gid,size,atime,mtime,ctime")
	for _, arg := range args {
		if _, err := os.Stat(arg); os.IsNotExist(err) {
			continue
		}
		// fmt.Println(i, arg)
		filepath.Walk(arg, stats)
	}
}