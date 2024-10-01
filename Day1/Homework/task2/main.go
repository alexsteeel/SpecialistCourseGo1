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

		reverseNumber := number%10*100 + number/10%10*10 + number/100
		fmt.Printf("Реверсная запись: %03d\n", reverseNumber)
	} else {
		fmt.Println("Число должно быть в диапазоне от -999 до 999 и из трех цифр.")
	}
}
