package main

import "fmt"

func main() {
	var number1, number2, number3 int
	fmt.Println("Введите три числа")
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	fmt.Scan(&number3)

	fmt.Println("Три числа, отсортированные в порядке возрастания:")
	if number1 < number2 && number1 < number3 {
		fmt.Print(number1)
		if number2 < number3 {
			fmt.Printf(" %d", number2)
			fmt.Printf(" %d", number3)
		} else {
			fmt.Printf(" %d", number3)
			fmt.Printf(" %d", number2)
		}
	} else if number2 < number1 && number2 < number3 {
		fmt.Print(number2)
		if number1 < number3 {
			fmt.Printf(" %d", number1)
			fmt.Printf(" %d", number3)
		} else {
			fmt.Printf(" %d", number3)
			fmt.Printf(" %d", number1)
		}
	} else {
		fmt.Print(number3)
		if number1 < number2 {
			fmt.Printf(" %d", number1)
			fmt.Printf(" %d", number2)
		} else {
			fmt.Printf(" %d", number2)
			fmt.Printf(" %d", number1)
		}
	}
}
