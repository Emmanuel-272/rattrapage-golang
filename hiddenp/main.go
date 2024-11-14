package main

import (
	"fmt"
	"os"
)

// func pour vérifié si s1 est cachée dans s2
func isHidden(s1, s2 string) bool {
	// si s1 est une chaîne vide, elle est considéré cmme caché dans n'importe quelle chaîne
	if s1 == "" {
		return true
	}

	// Parcourir les caractères de s2 et essayer de trouver chaque caractère de s1 dans le même ordre
	index := 0
	for i := 0; i < len(s2); i++ {
		if s2[i] == s1[index] {
			index++
			// Si tous les caractères de s1 sont trouvés, retourner true
			if index == len(s1) {
				return true
			}
		}
	}

	// Si tous les caractères de s1 sont pas été trouvés dans s2, retourner false
	return false
}

func main() {
	// apres on verifie si le nmbre d'argument est égal à 2
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	// on recuepere apres  les deux chaînes de caractères passées en argument
	s1 := os.Args[1]
	s2 := os.Args[2]

	// ensuite on  si s1 est caché dans s2
	if isHidden(s1, s2) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
