package main

import (
	"errors"
	"fmt"
)

const (
	min     = "minimum"
	average = "average"
	max     = "maximum"
)

func minim(nums ...float64) float64 {
	var mini float64 = -1
	for _, v := range nums {
		if mini == -1 {
			mini = v
		}
		if v < mini {
			mini = v
		}
	}
	return mini
}

func maxi(nums ...float64) float64 {
	var maxim float64 = 0
	for _, v := range nums {
		if v > maxim {
			maxim = v
		}
	}
	return maxim
}

func averag(nums ...float64) float64 {
	var average float64 = 0
	for _, v := range nums {
		average += v
	}
	return average / float64(len(nums))
}

func operacion(op string) (func(num ...float64) float64, error) {
	switch op {
	case min:
		return minim, nil
	case max:
		return maxi, nil
	case average:
		return averag, nil
	default:
		return nil, errors.New("Invalid operator")
	}
}

func main() {

	minFunc, _ := operacion(min)
	averFunc, _ := operacion(average)
	maxFunc, _ := operacion(max)

	fmt.Println(minFunc(1, 2, 3, 4))
	fmt.Println(averFunc(1, 2, 3, 4))
	fmt.Println(maxFunc(1, 2, 3, 4))
}
