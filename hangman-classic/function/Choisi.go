package hangman

import (
	"bufio"
	"math/rand"
	"os"
)

// ça permet de choisir un mot dans word.txt
func ChoisiMot(chemin string) (string, error) {
	file, err := os.Open(chemin) // initialise file au fichier
	if err != nil { // check les erreurs
		return "", err // retourne si erreur
	}
	defer file.Close() // ferme le file

	var mot []string // initialise mot
	scanner := bufio.NewScanner(file) // initialise scanner avec file
	for scanner.Scan() { // pour le scan
		mot = append(mot, scanner.Text()) // rajoute à mot le texte scanné
	}

	if err := scanner.Err(); err != nil { // check les erreurs
		return "", err // retourne si erreur
	}

	return mot[rand.Intn(len(mot))], nil // retourne le mot
}
