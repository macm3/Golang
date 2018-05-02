package main

import (
	"fmt"
	"os"

	. "github.com/macm3/Golang/one/src/functions"
)

func main() {
	fmt.Println("Escolha um comando:")
	ExibeMenu()
	var choice = ReadInt()

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
