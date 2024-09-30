package main

import (
	"fmt"
	"strings"
)

func main() {
	// Boolean => default:false
	var firstBoolean bool
	fmt.Println(firstBoolean)

	// Boolean operands
	aBoolean, bBoolean := true, false
	fmt.Println("AND:", aBoolean && bBoolean)
	fmt.Println("OR:", aBoolean || bBoolean)
	fmt.Println("NOT:", !aBoolean)

	// Классический условный оператор
	/*
		if condition1 {
			body1
		} else if condition2 {
			body2
		} else {
			body3
		}
	*/

	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Число", value, "четное")
	} else {
		fmt.Println("Число", value, "нечетное")
	}

	var color string
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknown color")
	}

	// Good Инициализация в блоке условного оператора
	// Блок присваивания -  только :=
	// Инициализируемая переменная видно только внутри области
	// видимости условного оператора (в телах if, else if или else),
	// но не за его пределами.
	var num int
	if fmt.Scan(&num); num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	// Не идиоматично
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}

	// Идиоматично
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
		return
	}
	fmt.Println("Width <= 100")
}
