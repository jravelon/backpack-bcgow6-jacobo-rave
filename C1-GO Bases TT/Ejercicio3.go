package main

import "fmt"

func main() {
	mes := 12
	mapa := map[int]string{1: "enero", 2: "febrero", 3: "marzo", 4: "abril",
		5: "mayo", 6: "junio", 7: "julio", 8: "agosto",
		9: "septiembre", 10: "octubre", 11: "noviembre", 12: "diciembre"}
	fmt.Println(mes, ", ", mapa[mes])
}
