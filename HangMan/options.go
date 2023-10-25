package main

import (
	"fmt"
	"os"
)

func Options() {
	fmt.Println("1) Начать Игру\n2) Выйти")

	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	switch input {
	case 1:
		fmt.Println("1")

	case 2:
		os.Exit(0)

	default:
		fmt.Println("Выберите цифру 1 или 2.")
		Options()
	}
}
