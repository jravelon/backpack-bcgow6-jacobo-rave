package main

import "fmt"

type User struct {
	Name    string
	Surname string
	Email   string
	Producs []Product
}

type Product struct {
	NameProduct     string
	PriceProduct    float64
	CuantityProduct int
}

func NewProduct(Name string, Price float64) *Product {
	newProd := Product{
		NameProduct:  Name,
		PriceProduct: Price,
	}
	return &newProd
}

func AddProduct(u *User, p *Product, Cuantity int) {
	p.CuantityProduct = Cuantity
	u.Producs = append(u.Producs, *p)
}

func DeleteProduct(u *User) {
	u.Producs = nil
}

func main() {
	u := &User{Name: "Juan", Surname: "Stive", Email: "js@gmail.com"}
	AddProduct(u, NewProduct("Cama", 500), 2)
	DeleteProduct(u)
	fmt.Print(u)
}
