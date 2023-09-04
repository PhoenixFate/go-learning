package main

import (
	"fmt"
	"os"
)

//go run 07-command-line.go  a b c
func main() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	for index, value := range os.Args {
		fmt.Printf("index: %d; value: %s\n", index, value)
	}

}
