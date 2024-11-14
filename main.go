package main

import (
	"fmt"
	"strings"
	"unicode"
)

// CapitalizeWords met en maj la premiere lettre et le reste en miniscule
func CapitalizeWords(s string) string {
	// Séparer la chaîne en mots
	words := strings.Fields(s)

	// il faut parcourir chaque mot pour modifier sa première lettre
	for i, word := range words {
		// Mettre la première lettre en majuscule et les autres en minuscule
		words[i] = string(unicode.ToUpper(rune(word[0]))) + strings.ToLower(word[1:])
	}

	// Rejoindre les mots en une seule chaîne et la retourner
	return strings.Join(words, " ")
}

func main() {
	original := "bonjour je suis un eleve d'ynov, comment allez vous, etc"
	result := CapitalizeWords(original)
	fmt.Println(result)
}
