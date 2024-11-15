package hangman

import (
	"math/rand"
	"strings"
)

// Cela permet de réveler des lettres aléatoirement
func RevLettres(mot string, lettreareveler int) string {
	revele := strings.Repeat("_", len(mot)) // initialise revele avec des _
	for i := 0; i < lettreareveler; i++ { // tant que i est inférieur à lettreareveler
		index := rand.Intn(len(mot)) // initialise index avec un chiffre random
		revele = revele[:index] + string(mot[index]) + revele[index+1:] // met revele avec la lettre sup
	}
	return revele // retourne revele
}