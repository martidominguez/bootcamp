package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	// handler
	handlerGreetings := func(w http.ResponseWriter, r *http.Request) {
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, "bad json decoding", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Hello " + person.FirstName + " " + person.LastName))

	}

	// endpoint
	http.HandleFunc("/greetings", handlerGreetings)

	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}

/* Ejercicio 2 - Manipulando el body


Vamos a crear un endpoint llamado /greetings.
Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hello + nombre + apellido”


El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hello Andrea Rivas”
La estructura deberá ser como esta:
{

                “firstName”: “Andrea”,

                “lastName”: “Rivas”

}
*/
