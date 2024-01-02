package food

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDogFood(t *testing.T) {
	dogAnimal := "dog"
	quantity := 10

	dogFunc, err := animal(dogAnimal)

	require.Equal(t, "", err)
	require.NotNil(t, dogFunc)
	require.Equal(t, 100, dogFunc(quantity))
}

func TestCatFood(t *testing.T) {
	catAnimal := "cat"
	quantity := 10

	catFunc, err := animal(catAnimal)

	require.Equal(t, "", err)
	require.NotNil(t, catFunc)
	require.Equal(t, 50, catFunc(quantity))
}

func TestHamsterFood(t *testing.T) {
	hamsterAnimal := "hamster"
	quantity := 10

	hamsterFunc, err := animal(hamsterAnimal)

	require.Equal(t, "", err)
	require.NotNil(t, hamsterFunc)
	require.Equal(t, 2500, hamsterFunc(quantity))
}

func TestTarantulaFood(t *testing.T) {
	tarantulaAnimal := "tarantula"
	quantity := 10

	tarantulaFunc, err := animal(tarantulaAnimal)

	require.Equal(t, "", err)
	require.NotNil(t, tarantulaFunc)
	require.Equal(t, 1500, tarantulaFunc(quantity))
}

func TestAnimalNotFound(t *testing.T) {
	animalNotFound := "animalNotFound"

	animalFunc, err := animal(animalNotFound)

	require.Equal(t, "Animal not found", err)
	require.Nil(t, animalFunc)
}
