// i have many doubts about this exercise, i don't know if i'm doing it right. i don't understand what i'm suppossed to do.

package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Home        string
}

var clients = []Client{
	{File: "file1.txt", Name: "Martina", ID: 1, PhoneNumber: "123456789", Home: "City1"},
	{File: "file2.txt", Name: "Juan", ID: 2, PhoneNumber: "987654321", Home: "City2"},
	{File: "file3.txt", Name: "Pedro", ID: 3, PhoneNumber: "123456789", Home: "City3"},
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	repeatedClient := Client{
		File:        "file1.txt",
		Name:        "Martina",
		ID:          1,
		PhoneNumber: "123456789",
		Home:        "City1",
	}

	nullClient := Client{
		File:        "",
		Name:        "John Doe",
		ID:          0,
		PhoneNumber: "1234567890",
		Home:        "123 Main Street",
	}

	validClient := Client{
		File:        "file4.txt",
		Name:        "John Doe",
		ID:          4,
		PhoneNumber: "1234567890",
		Home:        "123 Main Street",
	}

	registerNewClient(repeatedClient)

	registerNewClient(nullClient)

	registerNewClient(validClient)

	defer func() {
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()
}

func registerNewClient(client Client) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	err := checkNewClient(client)
	if err != nil {
		// option with panic
		panic(err.Error())

		// option with common errors
		// fmt.Println(err.Error())
	}

	err = checkNullData(client)
	if err != nil {
		// option with panic
		panic(err.Error())

		// option with common errors
		// fmt.Println(err.Error())
	}

	clients = append(clients, client)
	fmt.Printf("Client %s registered successfully\n", client.Name)
}

func checkNewClient(client Client) error {
	for _, c := range clients {
		if c.ID == client.ID {
			return errors.New("Error: client already exists")
		}
	}
	return nil
}

func checkNullData(client Client) error {
	if client.File == "" || client.Name == "" || client.ID == 0 || client.PhoneNumber == "" || client.Home == "" {
		return errors.New("Error: client data is null")
	}
	return nil
}

/*
Ejercicio 3 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos son:

File
Name
ID
Phone number
Home

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe.
Para ello, necesitás leer los datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto
hasta aquí. Ese error deberá:
1.- generar un panic;

2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar con la ejecución del programa normalmente.

Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos a
registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores.
Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes: “End of execution” y “Several errors were detected at runtime”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:

Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).
*/
