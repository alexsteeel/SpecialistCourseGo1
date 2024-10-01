package main

import "fmt"

func main() {
	var number int16
	fmt.Println("Введите четырехзначное число")
	fmt.Scan(&number)

	if (number >= -9999 && number < -999) || (number > 999 && number <= 9999) {
		if number < 0 {
			number = -number
		}

		firstTwoDigits := number / 100 % 100
		lastReversedTwoDigits := number%10*10 + number/10%10
		if firstTwoDigits == lastReversedTwoDigits {
			fmt.Printf("Число %d является палиндромом.", number)
			return
		}
		fmt.Printf("Число %d не является палиндромом.", number)
	} else {
		fmt.Println("Число должно быть в диапазоне от -9999 до 9999 и из четырех цифр.")
	}
}
