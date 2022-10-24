package calculadora

import "testing"

func TestSumar(t *testing.T) {
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	resultado := Sumar(num1, num2)
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma arrojo como resultado: %v, pero el esperado es: %v", resultado, resultadoEsperado)
	}
}
