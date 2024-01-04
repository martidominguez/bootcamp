package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "customers.txt"

	file, err := os.Open(fileName)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	defer file.Close()

	fmt.Println(string(fileContent))

	defer fmt.Println("ejecución finalizada")
}

/*
Ejercicio 2 - Imprimir datos


A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.

Ahora que el archivo sí existe, el panic no debe ser lanzado.

Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado,
debemos tener un “defer” que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.

es un type *os.File que te permite leerlo

para leerlo

io.Readall(file)

que te devuelve un slice de bytes y un err
bytes, err := io.ReadAll(file)
para transformarlo a string
string(bytes)
ya que se supone que los bytes que son un uint8 representan valores ascii, segurametne
por ende puede mapearlo
Lucas Nicolás Masiero17:31
tranquilamente a un string con string(bytes)
no funciona con todos
si los bytes son numeros random te va a aparece cuaquier cosa en la string
*/
