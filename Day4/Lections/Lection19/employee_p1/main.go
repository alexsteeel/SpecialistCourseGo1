package main

import (
	"fmt"
)

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// 1. Методы - функции, приаязанные к определенным структурам
// Шаблон
// func (s Struct) MethodName(parameters type) out type {}
// e в данном случае это получатель метода Receiver
func (e Employee) DisplayInfo() {
	fmt.Println("name", e.name)
	fmt.Println("position", e.position)
	fmt.Printf("salary %d %s\n", e.salary, e.currency)
}

func main() {
	// 2. Создание экземпляра
	emp := Employee{
		name:     "Mark",
		position: "Senior Golang programmer",
		salary:   500_000,
		currency: "rub",
	}

	emp.DisplayInfo()
}
