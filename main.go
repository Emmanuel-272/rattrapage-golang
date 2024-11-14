package main

import (
	"fmt"
	"os"
)

func main() {
	// Get the arguments passed to the program
	args := os.Args[1:]

	// Print the number of arguments
	fmt.Println(len(args))
}
