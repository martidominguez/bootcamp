package main

import "fmt"

func main() {
	var palabra string

	fmt.Print("Ingrese una palabra: ")
	fmt.Scanf("%s", &palabra)

	fmt.Printf("Cantidad de letras: %d\n", len(palabra))

	fmt.Println("Las letras de la palabra son:")

	for _, letra := range palabra {
		fmt.Printf("%c\n", letra)
	}
}
