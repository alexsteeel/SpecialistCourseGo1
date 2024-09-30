package main

import (
	"fmt"
	"math"
)

func main() {
	var result float64
	var d, e float64 = 10, 4

	var check float64 = 20 / 3
	a := 42
	b := 20
	c := a / b
	result = d / e

	fmt.Printf("type of c %T and value of c %v\n", c, c)
	fmt.Printf("type of result %T and value of result %v\n", result, result)
	fmt.Printf("type of check %T and value of check %v\n", check, check)

	fmt.Print("Инкремент: ")
	check++
	fmt.Println("check =", check)
	fmt.Print("Декремент: ")
	c--
	fmt.Println("c =", c)

	// Операция взятия остатка
	remainder := int(check) % 3
	fmt.Println("remainder:", remainder)

	// Операция возведения в степень
	var aa int = 2
	total := math.Pow(float64(aa), 10)
	fmt.Printf("total = %.2f\n", total)
}
