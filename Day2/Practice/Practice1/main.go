package main

import (
	"fmt"
)

func main() {
	var number uint32
	fmt.Println("Введите любое натуральное число.")
	fmt.Scan(&number)

	var sum uint32 = 0
	sumStr := ""
	digit := 0
	for i := 1; digit == 0; i++ {
		digit := number % 10
		number = number / 10
		sum += digit
		strDigit := fmt.Sprintf("%d", digit)

		if i == 1 {
			sumStr = strDigit
		} else {
			sumStr = sumStr + " + " + strDigit
		}
	}
	fmt.Printf("%v = %d\n", sumStr, sum)
}
