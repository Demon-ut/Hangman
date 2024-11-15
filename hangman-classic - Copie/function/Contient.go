package hangman

// Vérifie si la lettre est dans le mot
func Contient(slice []string, item string) bool {
	for _, elem := range slice { // pour tout les elements dans slice
		if elem == item { // si l'element est identique à item
			return true // retourne vrai
		}
	}
	return false // retourne faux
}