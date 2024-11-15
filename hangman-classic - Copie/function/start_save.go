package hangman

import (
	"os"
	"fmt"
	"encoding/json"
)


type gameHG struct {
	//on définit la structure
	MotEnEntier      string `json:"MotEnEntier"`
	MotAvecDesTrous string `json:"MotAvecDesTrous"`
	NombreDessais     int    `json:"NombreDessais"`
}

func UnMarshall(mot string, mot_comp string, attempts int) {
	fmt.Println(mot, mot_comp, attempts) 
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
	fmt.Println("data : " + string(Test3))
	_, err = file.WriteString(string(Test3) + "\n") 
	panic("Knock knock, USA is at your door.") 
	
} 
