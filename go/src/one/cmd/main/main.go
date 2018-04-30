package main

import (
	"fmt"
	"functions"
	"os"
)

func main() {
	fmt.Println("Escolha um comando:")
	functions.exibeMenu()
	var choice = readInt()

	if choice == 1 {

	} else if choice == 2 {

	} else if choice == 0 {
		fmt.Println("Tchau...")
		os.Exit(0)
	} else {
		fmt.Println("Escolha um comando que exista")
		os.Exit(-1)
	}

}
