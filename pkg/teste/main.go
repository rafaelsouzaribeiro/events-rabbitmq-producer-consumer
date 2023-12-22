package main

import "fmt"

func main() {
	eventos := []string{"teste", "teste1", "teste2"}
	//eventos = eventos[:2]
	eventos = append(eventos[:0], eventos[1:]...)

	// fmt.Println(eventos)

	fmt.Println(eventos)
}
