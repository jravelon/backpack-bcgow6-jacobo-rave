package main

import "fmt"

type Alumn interface {
	getName() string
	getLastName() string
	getDNI() string
	getDate() string
}

type Student struct {
	Name     string
	LastName string
	DNI      string
	Date     string
}

func (alumn Student) getName() string {
	return alumn.Name
}

func (alumn Student) getLastName() string {
	return alumn.LastName
}

func (alumn Student) getDNI() string {
	return alumn.DNI
}

func (alumn Student) getDate() string {
	return alumn.Date
}

func detalle(a Alumn) string {
	var infoStudent string
	infoStudent = ("Nombre: " + a.getName() + "\nApellido: " + a.getLastName() + "\nDNI: " + a.getDNI() + "\nFecha: " + a.getDate())
	return infoStudent
}

func main() {
	student := Student{
		Name:     "John",
		LastName: "Doe",
		DNI:      "10",
		Date:     "14 julio 2022",
	}
	fmt.Println(detalle(student))
}
