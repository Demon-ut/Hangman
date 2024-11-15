package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fonction General() qui lance le jeu du pendu
func General() {
	var Reset = "\033[0m"    // initialise Reset
	var Red = "\033[31m"     // initialise Red
	var Green = "\033[32m"   // initialise Green
	var Yellow = "\033[33m"  // initialise Yellow
	var Blue = "\033[34m"    // initialise Blue
	var Magenta = "\033[35m" // initialise Magenta
	var Cyan = "\033[36m"    // initialise Cyan
	//var Gray = "\033[37m" // initialise Gray
	var White = "\033[97m" // initialise White
	score := 0             // initialise score à 0
	var rep string         // initialise rep en string
	for rep != "q" {       // démarre une boucle de répétition
		print(Red + "T" + Reset)     // affiche T
		print(Green + "y" + Reset)   // affiche y
		print(Yellow + "p" + Reset)  // affiche p
		print(Blue + "e" + Reset)    // affiche e
		print(" ")                   // affiche un espace
		print(Magenta + "s" + Reset) // affiche sm
		print(Cyan + "o" + Reset)    // affiche o
		print(Red + "m" + Reset)     // affiche m
		print(Green + "e" + Reset)   // affiche e
		print(Yellow + "t" + Reset)  // affiche t
		print(Blue + "h" + Reset)    // affiche h
		print(Magenta + "i" + Reset) // affiche i
		print(Cyan + "n" + Reset)    // affiche n
		print(Red + "g" + Reset)     // affiche g
		print(Green + ":" + Reset)   // affiche :
		fmt.Scanln(&rep)             // Attend une réponse
		if rep != "q" {              // Vérifie que la réponse est q

			// Choisir un mot
			mot, err := ChoisiMot("words.txt") // words.txt
			if err != nil {                    // vérifie qu'il n'y a pas d'erreur
				fmt.Println("Error reading file:", err) // affiche l'erreur si présente
				return
			}

			// Initialiser les essais et lettres à révéler
			attempts := 10                       // initialise le nombre d'essais
			lettreareveler := (len(mot) / 2) - 1 // décide du nombre de lettres à réveler dans le mot
			if lettreareveler < 1 {              // vérifie que le nombre de lettres à réveler est inférieur à 1
				lettreareveler = 1 // met le nombre de lettres à réveler à 1
			}

			revele := RevLettres(mot, lettreareveler)                       // révele des lettres aléatoires (dépendant de lettreareveler)
			fmt.Println(Green + "Good luck, you have 10 attempts." + Reset) // affiche Good luck, you have 10 attempts.
			ChangeValue(revele)                                             // affiche le mot en ASCII art
			var lettredevine []string                                       // initialise lettredevine en tableau

			// Boucle principale pour gérer les tentatives
			for attempts > 0 && strings.Contains(revele, "_") { // boucle tant que attempts est supérieur à 0 et que il reste des lettres non découvertes
				fmt.Print("Choose: ")                         // affiche Choose:
				lecture := bufio.NewReader(os.Stdin)          // ???
				ans, _ := lecture.ReadString('\n')            // ???
				ans = strings.TrimSpace(strings.ToLower(ans)) // met la réponse en minuscule pour éviter les erreurs

				//Stop le jeu
				if ans == "stop" { // vérifie que la réponse soit stop
					Marshall(revele, mot, attempts) // sauvegarde la partie dans save.txt
				}

				//Secret
				if ans == "Mario" || ans == "mario" { // vérifie que la réponse soit mario
					Mario() // affiche mario
				}

				// Vérifier si l'utilisateur met une seule lettre
				if len(ans) > 1 || !strings.ContainsAny(ans, "abcdefghijklmnopqrstuvwxyz") { // vérifie que la longueur de la réponse dépasse 1 ou que la réponse n'est pas une lettre
					fmt.Println(Red + "Please enter a single letter." + Reset) // affiche Please enter a single letter
					fmt.Println()                                              // affiche un retour à la ligne
				}

				// Si l'utilisateur a déjà deviné la lettre
				if Contient(lettredevine, ans) { // vérifie que la liste de lettres devinée contient la réponse
					fmt.Println(White + "You've already guessed that letter.") // affiche You've already guessed that letter.
					fmt.Println()                                              // affiche un retour à la ligne
					continue                                                   // continue
				} else { // sinon
					lettredevine = append(lettredevine, ans) // rajoute la lettre à la liste de lettres devinée
					lettredevine = append(lettredevine, " || ")
				}

				// Vérifier si la lettre devinée est correcte
				if strings.Contains(mot, ans) { // vérifie que le mot possède la réponse
					for i, letter := range mot { // boucle pour chaque lettres dans mot
						if string(letter) == ans { // vérifie que la lettre soit identique à la réponse
							revele = revele[:i] + string(letter) + revele[i+1:] // révele la lettre
						}
					}
					fmt.Println(Green + "Good guess!" + Reset) // affiche Good guess!
				} else { // sinon
					fmt.Println(Pendu(attempts))                                                        // affiche le pendu (dépend des attempts restant)
					attempts--                                                                          // retire 1 à attempts
					fmt.Printf(Red+"Not present in the word, remaining attempts: %d\n"+Reset, attempts) // affiche Not present in the world, remaining attempts: et les attempts restant
				}

				// Condition de victoire
				if mot == ans || revele == mot { // vérifie que la réponse soit le mot
					fmt.Println(Green + "Congrats! You guessed the whole word!" + Reset) // affiche Congrats! You guessed the whole word!
					PrintWin()                                                           // affiche la victoire
					ChangeValue(mot)                                                     // affiche le mot en ASCII
					score += 1                                                           // ajoute 1 à score
					break                                                                // arrête la boucle principale
				} else if mot != ans && len(ans) > 1 { // sinon vérifie que mot est différent de la réponse et que la longueur de la réponse dépasse 1
					attempts -= 2 // retire 2 à attempts
					continue      // continue
				}

				//Condition perdu
				if attempts<=0 { // 
					PrintLose() // 
					fmt.Println("You lost! The word was:") // 
					ChangeValue(mot) // 
					fmt.Println("Final score: ", score) // 
					score = 0 // 
					continue // 
				}

				fmt.Println()                 // affiche un retour à la ligne
				fmt.Println("Word to guess:") // affiche "Word to guess"
				ChangeValue(revele)           // affiche Word to guess et le mot semi caché
				fmt.Println("Letters already used: ", lettredevine) // 
			}
		}
	}
	fmt.Printf("Final score: %d\n", score) // affiche le score final
}
