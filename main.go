package main

import (
	"fmt"
	"go-learning/testModule"
	"io/fs"
	"os"
)

func main() {
	var testString string = "test"
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
	test(&testString)
	fmt.Println(testString)
	testModule.TestFunction()
	testFor()
}

func test(testString *string) {
	*testString = "change"
	fmt.Println(*testString)
	test2(testString)
}

func test2(testString *string) {
	*testString = "change2"
	fmt.Println(*testString)
}

func testFor() {
	var slice = []string{"hello", "this", "is a", "test"} //this is called a slice

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	for _, elem := range slice {
		fmt.Println(elem)
	}
}
