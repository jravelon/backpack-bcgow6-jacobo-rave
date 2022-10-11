package main

import (
	"os"
)

func main() {
	producto := "id;price;quanity\n"
	producto += ("1;11;123\n")
	producto += ("2;22;1\n")

	os.WriteFile("product.csv", []byte(producto), 0644)
}
