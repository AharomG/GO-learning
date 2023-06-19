package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	var file *os.File
	var err error
	file, err = os.Open("values.txt")
	var info fs.FileInfo
	info, err = file.Stat()
	var size int64 = info.Size()
	var numberOfReadBytes int
	var buf []byte = make([]byte, size)
	fmt.Println(size)
	fmt.Println(file)
	fmt.Println(err)
	numberOfReadBytes, err = file.Read(buf)
	fmt.Println(err)
	fmt.Println(numberOfReadBytes)
	fmt.Println(string(buf))
}
