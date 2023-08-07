package main

import (
	"io"
	"os"
)

func main() {
	filename := getFileName()
	readFile(filename)
}

func getFileName() string {
	filename := os.Args[1]
	return filename
}

func readFile(filename string) {
	println("Read file:" + filename)
	dat, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, dat)
}
