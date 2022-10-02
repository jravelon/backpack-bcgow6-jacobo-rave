package main

import (
	"errors"
	"fmt"
)

const (
	dog     = "dog"
	cat     = "cat"
	hamster = "hamster"
	spider  = "spider"
)

func feedDog(cantity float64) float64 {
	return float64(cantity) * 10
}
func feedCat(cantity float64) float64 {
	return float64(cantity) * 5
}
func feedHamster(cantity float64) float64 {
	return float64(cantity) * 0.25
}
func feedSpider(cantity float64) float64 {
	return float64(cantity) * 0.15
}

func tipoAnimal(name string) (func(cantity float64) float64, error) {
	switch name {
	case dog:
		return feedDog, nil
	case cat:
		return feedCat, nil
	case hamster:
		return feedHamster, nil
	case spider:
		return feedSpider, nil
	default:
		return nil, errors.New("Animal " + name + " don't exist")
	}
}

func tipoAnimalCantity(name string, n float64) (float64, error) {
	animal, err := tipoAnimal(name)
	if err == nil {
		return animal(n), nil
	} else {
		fmt.Println("Animal " + name + " don't exist")
	}
	return 0, nil
}

func main() {
	var totalCantity float64 = 0
	cantity1, _ := tipoAnimalCantity(dog, 5)
	cantity2, _ := tipoAnimalCantity(cat, 4)
	cantity3, _ := tipoAnimalCantity(hamster, 4)
	cantity4, _ := tipoAnimalCantity(spider, 1)
	cantity5, _ := tipoAnimalCantity("pajaro", 5)
	totalCantity += cantity1 + cantity2 + cantity3 + cantity4 + cantity5
	fmt.Println("El total de Kg de comida a comprar es: ", totalCantity)

}
