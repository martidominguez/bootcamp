package main

import "fmt"

type Small struct {
	InitialPrice float64
}

type Medium struct {
	InitialPrice float64
}

type Large struct {
	InitialPrice float64
}

type Product interface {
	Price() float64
}

func (s Small) Price() float64 {
	return s.InitialPrice
}

func (m Medium) Price() float64 {
	return m.InitialPrice * 1.03
}

func (l Large) Price() float64 {
	return l.InitialPrice*1.06 + 2500
}

func factory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return Small{
			InitialPrice: price,
		}
	case "Medium":
		return Medium{
			InitialPrice: price,
		}
	case "Large":
		return Large{
			InitialPrice: price,
		}
	default:
		return nil
	}
}

func main() {

	product := factory("Large", 10000)

	fmt.Println(product.Price())

	product = factory("Medium", 10000)

	fmt.Println(product.Price())

	product = factory("Small", 10000)

	fmt.Println(product.Price())
}
