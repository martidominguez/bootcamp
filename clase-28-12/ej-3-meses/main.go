package main

import "fmt"

func main() {
	var mesSeleccionado int
	meses := [12]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}

	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&mesSeleccionado)

	if mesSeleccionado < 1 || mesSeleccionado > 12 {
		fmt.Println("No se ingreso un mes valido")
	} else {
		fmt.Printf("%d, %s\n", mesSeleccionado, meses[mesSeleccionado-1])
	}
}

/* Para resolverlo de otra forma, se podría realizar un switch. Sin embargo, considero que de esta forma es mas eficiente. */
