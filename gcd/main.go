package main

import (
	"fmt"
	"os"
	"strconv"
)

// GCD calcule le plus grand commun diviseur de deux entiers
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func main() {
	// je verifie si le nombre d'arguments est égal à 2
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// et je converti donc les arguments en entier
	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	// apres je verifie si la conversion a réussi
	if err1 != nil || err2 != nil {
		fmt.Println()
		return
	}

	// je calcul et ensuite j'affiche le PGCD
	fmt.Println(GCD(num1, num2))
}
