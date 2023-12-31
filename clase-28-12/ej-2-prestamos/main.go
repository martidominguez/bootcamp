package main

import "fmt"

func main() {
	var edad int
	var empleado bool
	var antiguedad int = 0
	var sueldo int = 0

	fmt.Printf("Ingrese su edad: ")
	fmt.Scanln(&edad)

	fmt.Printf("Se encuentra empleado?(true/false) ")
	fmt.Scanln(&empleado)

	if empleado {
		fmt.Printf("Años de antiguedad: ")
		fmt.Scanln(&antiguedad)

		fmt.Printf("Sueldo: ")
		fmt.Scanln(&sueldo)
	}

	switch {
	case edad > 22 && empleado && antiguedad > 1 && sueldo > 100000:
		fmt.Printf("Se le otorgara un prestamo sin intereses\n")
	case edad > 22 && empleado && antiguedad > 1:
		fmt.Printf("Se le otorgara un prestamo con intereses\n")
	default:
		fmt.Printf("No se le entregara un prestamo\n")
	}

	/* if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Printf("Se le otorgara un prestamo sin intereses\n")
		} else {
			fmt.Printf("Se le otorgara un prestamo con intereses\n")
		}
	} else {
		fmt.Printf("No se le entregara un prestamo\n")
	} */

}
