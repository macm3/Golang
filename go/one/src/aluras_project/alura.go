package main

import "fmt"
import "os"

func main() {
	fmt.Println("Escolha um comando:")
	exibeMenu()
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

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readInt() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
