package main

import (
	"log"

	"github.com/clukawski/wordle"
)

func main() {
	we, err := NewWordleEngine("./dictionary", 8)
	if err != nil {
		log.Fatalln(err)
	}

	we.NewGame()
	log.Printf("Answer: %s\n", string(we.CurrentGame.Answer))

	for i := 0; i < we.MaxAttempts; i++ {
		guess, err := we.RandomDictionaryWord()
		if err != nil {
			log.Println("Failed to get random dictionary word")
			break
		}

		correct, err := we.CurrentGame.Attempt(guess)
		if err != nil {
			log.Printf("Game over! The word was: %s", we.CurrentGame.GetAnswer())
			break
		}
		if correct {
			log.Printf("You win! The word was: %s", we.CurrentGame.GetAnswer())
			break
		}
	}
	for _, attempt := range we.CurrentGame.Attempts {
		log.Printf("Attempt: %+v\n", attempt)
	}
}
