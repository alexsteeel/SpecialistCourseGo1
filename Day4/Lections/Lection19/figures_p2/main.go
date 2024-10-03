package main

import (
	"fmt"
	"math"
)

// 4. В чем преимущество методов над функциями
// Во-первых: наличие методов улучшает читаемость кода (консистентность кода)
// Во-вторых: в go запрещается создание функций с одинаковыми названиями,
// а методы для различных структур могут иметь одинаковые имена

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width int
}

// 4. Создадим функцию и метод Perimeter для структуры Circle
func PerimeterCircle(c Circle) float64 {
	return 2 * c.radius * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * c.radius * math.Pi
}

// 4. Создадим функцию и метод Perimeter для структуры Rectangle
func PerimeterRectangle(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("Call function for circle:", PerimeterCircle(c))
	fmt.Println("Call method for circle:", c.Perimeter())

	r := Rectangle{10, 5}
	fmt.Println("Call function for rectangle:", PerimeterRectangle(r))
	fmt.Println("Call method for rectangle:", r.Perimeter())
}
