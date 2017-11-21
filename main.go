package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var count int64
var inputFile string

func main() {
	flag.StringVar(&inputFile, "f", "", "file name")
	flag.Parse()
	if inputFile == "" {
		fmt.Println("Use -f <filename> to specify a filename")
		os.Exit(1)
	}
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
	}
	for _, j := range string(buf) {
		// reference https://jimmysong.io/cheatsheets/unicode
		if int64(j) > hex2oct("0x4e00") && int64(j) < hex2oct("0x9fa5") {
			count++
		}
	}
	fmt.Println(inputFile, count)
}

func hex2oct(hex string) int64 {
	oct, _ := strconv.ParseInt(hex, 0, 64)
	return oct
}
