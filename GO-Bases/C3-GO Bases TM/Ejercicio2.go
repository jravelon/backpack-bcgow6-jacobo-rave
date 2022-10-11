package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	table, err := os.ReadFile("product.csv")
	tableStr := string(table)
	a := strings.Split(tableStr, "\n")
	aJoin := strings.Join(a, " ")
	b := strings.Split(aJoin, ";")
	bJoin := strings.Join(b, " ")
	splited := strings.Split(bJoin, " ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("-------------------------------------------------------\n")
	fmt.Printf("| %-10v | %10v | %10v |\n", splited[0], splited[1], splited[2])
	fmt.Printf("|------------|------------|------------|\n")

	var arr []string = splited[3:]
	sum := 0
	j := 0
	for i := range arr {

		if j+1 > len(arr) || j+2 > len(arr) || j+3 > len(arr) {
			break
		}
		fmt.Printf("| %-10v | %10v | %10v |\n", arr[j], arr[j+1], arr[j+2])

		num, _ := strconv.Atoi(arr[j+1])
		sum += num

		i++
		j = i * 3

	}
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("  %10v : %10v \n", "Total", sum)
}
