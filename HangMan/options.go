package main

import (
	"fmt"
	"os"
)

func Options() {
	fmt.Println("1) Start Game\n2) Exit")

	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	switch input {
	case 1:
		StartGame()

	case 2:
		os.Exit(0)

	default:
		fmt.Println("Select number 1 or 2.")
		Options()
	}
}
