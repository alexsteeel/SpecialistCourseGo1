package main

import "fmt"

func main() {
	var number int16
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)

	if number >= 1000 || number < -999 {
		fmt.Println("Число должно быть в диапазоне от -999 до 999")
	} else {
		if number < 0 {
			number = -number
		}

		fmt.Println("Первая цифра:", number/100)
		fmt.Println("Вторая цифра:", number/10%10)
		fmt.Println("Третья цифра:", number%10)
	}
}
