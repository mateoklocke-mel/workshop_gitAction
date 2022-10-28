package main

import "fmt"

func main() {
	fmt.Println("Hola esto es una prueba")
	fmt.Println("suma:", suma(4, 4))
}

func suma(a, b int) int {
	return a + b
	//return a - b
}
