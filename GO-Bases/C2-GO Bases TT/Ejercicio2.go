package main

import "fmt"

type Matrix struct {
	width     int
	height    int
	cuadratic bool
	maxValue  float64
	matrix    [][]float64
}

func (m Matrix) set(values ...float64) {
	k := 0
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			if k == len(values) {
				break
			}
			m.matrix[i][j] = values[k]
			k++
		}
	}
}

func createMatrix(width int, height int) Matrix {
	matrixToCreate := Matrix{width: width, height: height}
	matrixToCreate.matrix = make([][]float64, width)
	for i := range matrixToCreate.matrix {
		matrixToCreate.matrix[i] = make([]float64, height)
	}
	if width == height {
		matrixToCreate.cuadratic = true
	}
	return matrixToCreate
}

func (m Matrix) printMatrix() {
	for _, value := range m.matrix {
		fmt.Println(value)
	}
}

func main() {
	newMatrix := createMatrix(3, 3)
	newMatrix.printMatrix()
	newMatrix.set(1, 2, 3, 4, 5, 6, 7, 8, 9)
	newMatrix.printMatrix()
}
