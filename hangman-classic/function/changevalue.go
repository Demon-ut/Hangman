package hangman

import (
	"os"
	"bufio"
	"fmt"
)

func ChangeValue(mot string) {
	var tab[27][6]string
	file, err := os.Open("standard.txt") // initialise file au fichier
	if err != nil { // check les erreurs
		return // retourne si erreur
	}
	defer file.Close() // ferme le file
	scanner := bufio.NewScanner(file) // scan le fichier
	for i:=0; i<27; i++ { // tant que i est inférieur à 27
		for j:=0; j<6; j++ { // tant que j est inférieur à 6
			if scanner.Scan() { // si l'on récup des infos du fichier
				tab[i][j] = scanner.Text() // rajoute les infos dans un tableau
			}
		}
	}
	for lineIndex:=0; lineIndex<6; lineIndex++{ // tant que l'index est inférieur à 6
		for _, elmt := range mot{ // pour tout les elmts dans mot
			if elmt == '_' { // si l'élement est _
				fmt.Print(tab[elmt-69][lineIndex]) // converti en Ascii de _
			} else { // sinon
			fmt.Print(tab[elmt-97][lineIndex]) // converti en Ascii de la lettre
		}}
		fmt.Println() // retour à la ligne
	}
}
