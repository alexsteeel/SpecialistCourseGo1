package main

import (
	"fmt"
)

func main() {
	var number uint32
	fmt.Println("Введите любое натуральное число.")
	fmt.Scan(&number)

	var sum uint32 = 0
	statement := ""
	for number != 0 {
		digit := number % 10
		number /= 10
		sum += digit
		strDigit := fmt.Sprintf("%d", digit)

		if len(statement) == 0 {
			statement = strDigit
		} else {
			statement = statement + " + " + strDigit
		}
	}
	fmt.Printf("%v = %d\n", statement, sum)
}
