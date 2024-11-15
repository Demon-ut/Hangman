package hangman

import (
	"bufio"
	"os"
)

func Pendu(attempts int) string {
	res := "" // initialise le résultat
	file, _ := os.Open("hangman.txt") // initialise file avec le fichier
	defer file.Close() // ferme le fichier

	var pendu []string // initialise pendu
	scanner := bufio.NewScanner(file) // initialise scanner avec file
	for scanner.Scan() { // pour le contenu de scanner
		pendu = append(pendu, scanner.Text()) // rajoute le texte de scanner dans pendu
	}
	// POUR TOUT LE RESTE
	// si attemps == le chiffre
	// tant que i est inférieur à (la valeur mise)
	// ajoute l'état suivant du pendu
	if attempts == 10 {
		for i := 0; i <= 7; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 9 {
		for i := 8; i <= 15; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 8 {
		for i := 16; i <= 23; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 7 {
		for i := 24; i <= 31; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 6 {
		for i := 32; i <= 39; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 5 {
		for i := 40; i <= 47; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 4 {
		for i := 48; i <= 55; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 3 {
		for i := 56; i <= 63; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 2 {
		for i := 64; i <= 71; i++ {
			res += pendu[i] + "\n"
		}
	}
	if attempts == 1 {
		for i := 72; i <= 78; i++ {
			res += pendu[i] + "\n"
		}
	}
	return res // affiche le pendu
}