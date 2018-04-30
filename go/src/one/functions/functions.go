package functions

import (
	"fmt"
)

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
