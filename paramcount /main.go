package main

import (
	"fmt"
	"os"
)

func main() {
	// On obtient les arguments transmis au prgramme
	args := os.Args[1:]

	// on imprime le nmbre d'arguments
	fmt.Println(len(args))
}
