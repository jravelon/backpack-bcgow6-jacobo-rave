package main

import "fmt"

const (
	big    = "big"
	medium = "medium"
	small  = "small"
)

type ProductInterface interface {
	CalcCost() float64
}

type Ecommerce interface {
	Total() float64
	Add(p ProductInterface)
	Printin()
}

type Store struct {
	ListProducts []ProductInterface
}

type Product struct {
	TypeProduct  string
	NameProduct  string
	PriceProduct float64
}

func NewProduct(TypeProduct, NameProduct string, PriceProduct float64) *Product {
	NewProducto := Product{
		TypeProduct:  TypeProduct,
		NameProduct:  NameProduct,
		PriceProduct: PriceProduct,
	}
	return &NewProducto
}

func NewStore() Ecommerce {
	return &Store{}
}

func (p Product) CalcCost() float64 {
	switch p.TypeProduct {
	case small:
		return p.PriceProduct
	case medium:
		return p.PriceProduct + (p.PriceProduct * 0.03)
	case big:
		return p.PriceProduct + (p.PriceProduct * 0.06) + 2500
	}
	return 0
}

func (s Store) Total() float64 {
	var total float64 = 0
	for _, value := range s.ListProducts {
		total += value.CalcCost()
	}
	return total
}

func (s *Store) Add(p ProductInterface) {
	s.ListProducts = append(s.ListProducts, p)
}

func (s *Store) Printin() {
	for _, value := range s.ListProducts {
		fmt.Printf("%v\n", value)
	}
}

func main() {
	e := NewStore()
	e.Add(NewProduct(big, "a", 10))
	e.Add(NewProduct(big, "b", 10))
	e.Add(NewProduct(big, "c", 10))
	fmt.Println(e.Total())
	e.Printin()
}
