package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(os.Args)

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", filename)
		os.Exit(1)
	}

	numWritten, err := io.Copy(os.Stdout, file)

	if err != nil {
		fmt.Println("Error reading file to stdout.")
		os.Exit(1)
	}
	fmt.Println("\nFile finished being printed.", numWritten, "bytes read.")
}
