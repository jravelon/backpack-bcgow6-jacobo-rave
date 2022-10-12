/*Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la
velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el
total de los 3).
*/

package main

import "fmt"

func main() {
	products := []Products{
		{NameP: "p1", PriceP: 100, CuantityP: 10},
		{NameP: "p2", PriceP: 200, CuantityP: 10},
		{NameP: "p3", PriceP: 3, CuantityP: 10},
	}

	services := []Services{
		{NameS: "s1", PriceS: 100, Minutes: 120},
		{NameS: "s1", PriceS: 300, Minutes: 15},
	}

	mantainences := []Maintenance{
		{NameM: "m1", PriceM: 1000},
		{NameM: "m2", PriceM: 3000},
		{NameM: "m3", PriceM: 6000},
	}

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go sumProducts(products, c1)
	go sumMaintenance(mantainences, c2)
	go sumServices(services, c3)

	for i := 0; i < 3; i++ {
		select {
		case ans1 := <-c1:
			fmt.Println("Products sum: ", ans1)
		case ans2 := <-c2:
			fmt.Println("Mantainences sum: ", ans2)
		case ans3 := <-c3:
			fmt.Println("Services sum: ", ans3)
		}
	}
}

func sumProducts(arr []Products, c chan int) {
	sum := 0.0
	for _, value := range arr {
		sum += value.PriceP * value.CuantityP
	}
	c <- int(sum)

}

func sumServices(arr []Services, c chan int) {
	sum := 0
	for _, value := range arr {
		sum += value.PriceS / 2 * ((value.Minutes + 29) / 30)
	}
	c <- sum
}

func sumMaintenance(arr []Maintenance, c chan int) {
	sum := 0
	for _, value := range arr {
		sum += value.PriceM
	}
	c <- sum
}

type Products struct {
	NameP     string
	PriceP    float64
	CuantityP float64
}

type Services struct {
	NameS   string
	PriceS  int
	Minutes int
}

type Maintenance struct {
	NameM  string
	PriceM int
}
