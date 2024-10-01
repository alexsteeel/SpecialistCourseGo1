package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Println("Введите три числа")
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	fmt.Scan(&number3)

	if number1 > number2 {
		number1, number2 = number2, number1
	}

	if number2 > number3 {
		number2, number3 = number3, number2
	}

	if number1 > number2 {
		number1, number2 = number2, number1
	}

	fmt.Printf("Три числа, отсортированные в порядке возрастания: %d %d %d\n", number1, number2, number3)
}
