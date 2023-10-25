package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func StartGame() {
	n := rand.Intn(len(Words))

	randomWord := Words[n]
	displayWord := strings.Repeat("_", len(randomWord))

	i := 0
	for {
		fmt.Println(HangmanPositions[i])

		if strings.EqualFold(displayWord, randomWord) {
			fmt.Printf("Congratulations, you won the game! The word was %q\n", randomWord)
			break
		}

		if i == len(HangmanPositions)-1 {
			fmt.Println("You losed :(")
			break
		}

		fmt.Printf("Word to guess: %s\n", displayWord)

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(input) != 1 {
			fmt.Println("Please enter a single letter. Try again.")
			continue
		}

		if strings.Contains(randomWord, input) {
			for j, letter := range randomWord {
				if strings.EqualFold(string(letter), input) {
					displayWord = displayWord[:j] + string(letter) + displayWord[j+1:]
				}
			}
		} else {
			i++
		}
	}
}
