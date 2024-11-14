package main

import (
	"fmt"
	"os"
	"strconv"
)

// une func pour vérifié si un nmbre est premier
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// une func qui calcul la smme des nmbres premier inférieurs ou égal à un nmbre donné
func sumPrimes(limit int) int {
	sum := 0
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

func main() {
	// une vérification du nmbre d'argument
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	// on converti l'argument en entier
	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 {
		fmt.Println(0)
		return
	}

	// on calcul et on affiche la smme des nmbres premiés
	fmt.Println(sumPrimes(num))
}
