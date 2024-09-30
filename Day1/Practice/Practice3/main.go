package main

import "fmt"

func main() {
	var number int16
	fmt.Println("Введите трехзначное число")
	fmt.Scan(&number)

	if (number >= -999 && number < -99) || (number > 99 && number <= 999) {
		if number < 0 {
			number = -number
		}

		fmt.Println("Первая цифра:", number/100)
		fmt.Println("Вторая цифра:", number/10%10)
		fmt.Println("Третья цифра:", number%10)
	} else {
		fmt.Println("Число должно быть в диапазоне от -999 до 999 и из трех цифр.")
	}
}
