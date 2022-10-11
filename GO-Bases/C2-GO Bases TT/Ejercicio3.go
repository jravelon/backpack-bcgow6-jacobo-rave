/*Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
-Crear una estructura “tienda” que guarde una lista de productos.
-Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
-Crear una interface “Producto” que tenga el método “CalcularCosto”
-Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
-Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
-Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
-El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

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
