package food

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func main() {
	dogAnimal := dog
	quantity := 10

	dogFunc, err := animal(dogAnimal)

	if err != "" {
		fmt.Printf("Error: %s\n", err)
		return
	}

	amountOfFood := dogFunc(quantity)

	fmt.Printf("Amount of food for %d %ss: %d\n", quantity, dogAnimal, amountOfFood)

}

func animal(animal string) (animalFunc func(int) int, err string) {
	switch animal {
	case "dog":
		animalFunc = animalDog
	case "cat":
		animalFunc = animalCat
	case "hamster":
		animalFunc = animalHamster
	case "tarantula":
		animalFunc = animalTarantula
	default:
		err = "Animal not found"
	}
	return
}

func animalDog(quantity int) int {
	return quantity * 10
}

func animalCat(quantity int) int {
	return quantity * 5
}

func animalHamster(quantity int) int {
	return quantity * 250
}

func animalTarantula(quantity int) int {
	return quantity * 150
}

/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.

Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y
que retorne una función y un mensaje (en caso que no exista el animal).
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
Ejemplo:

const (
   dog    = "dog"
   cat    = "cat"
)
...
animalDog, msg := animal(dog)
animalCat, msg := animal(cat)
...
var amount float64
amount += animalDog(10)
amount += animalCat(10)
*/
