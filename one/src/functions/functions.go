package functions

import (
	"fmt"
)

func ExibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func ReadInt() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}
