package main

import (
	"fmt"
)

// RecursivePower calcule la puissance d'un entier de manière récursive
// si l'exposant est négatif il retourne 0
func RecursivePower(base, exponent int) int {
	// si l'exposant est négatif retourner 0
	if exponent < 0 {
		return 0
	}

	// si l'exposant est 0, retourner 1
	if exponent == 0 {
		return 1
	}

	// dans le cas récursif il faut multiplier la base par la puissance avec un exposant réduit de 1
	return base * RecursivePower(base, exponent-1)
}

func main() {
	// un exemple assez simple
	fmt.Println(RecursivePower(2, 3))  // doit afficher 8
	fmt.Println(RecursivePower(5, -2)) // doit afficher 0
	fmt.Println(RecursivePower(3, 0))  // doit afficher 1
}
