package main

import "fmt"

func main() {
	fmt.Println("Hello Word!")
	fmt.Println("CHANGING TO COMPILE AND RUN AT THE SAME TIME")
	name := "Marcela"
	lastName := "Azevedo"
	age := 22
	//fmt.Println(reflect.TypeOf(name))
	fmt.Println(name, lastName+".", "I'm", age, "years old")
	fmt.Println("Write two numbers at the same line so we can add them")
	var first int
	var second int
	fmt.Scan(&first, &second)
	var result = plus(first, second)
	fmt.Println("The result is:", result)
}

func Plus(a int, b int) int {
	return a + b
}
