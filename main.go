package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// get filename from Args[1]
	filename := os.Args[1]
	// read file
	file, foerr := os.Open(filename)
	// log errors and exit
	if foerr != nil {
		log.Fatal("Error opening file:", foerr)
	}
	// io.Copy to stdout
	io.Copy(os.Stdout, file)
}
