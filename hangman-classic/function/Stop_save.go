package hangman

import (
	"encoding/json"
	"fmt"
	"os"
)


type game struct {
	//on définit la structure
	MotEnEntier      string `json:"MotEnEntier"`
	MotAvecDesTrous string `json:"MotAvecDesTrous"`
	NombreDessais     int    `json:"NombreDessais"`
}

func Marshall(mot string, mot_comp string, attempts int) {
	file, err := os.OpenFile("save.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) //on n'ouvre le file
	defer file.Close()//on le ferme
	
	encoding := game{ //on encode
		MotEnEntier:      mot,
		MotAvecDesTrous: mot_comp,
		NombreDessais:     attempts,
	}
	Test3, err := json.Marshal(encoding) 
	if err != nil {                      
		fmt.Println("Error:", err) 
		return
	}
	_, err = file.WriteString(string(Test3) + "\n") 
	fmt.Println("votre partie à bien été sauvegardé")
	os.Exit(0)
} 
