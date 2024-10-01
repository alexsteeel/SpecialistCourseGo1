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

		if len(statement) == 0 {
			statement = fmt.Sprintf("%d", digit)
		} else {
			statement = statement + " + " + fmt.Sprintf("%d", digit)
		}
	}
	fmt.Printf("%v = %d\n", statement, sum)
}
