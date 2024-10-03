package main

import (
	"fmt"
)

type Employee struct {
	name   string
	salary int
}

// 1. Метод, в котором получатель копируется и в его теле происходит работа с локаьлной копией
func (e Employee) SetName(newName string) {
	e.name = newName
}

// 2. Метод, в котором получатель передается по ссылке (в теле метода работаем со ссылкой на экземпляр)
func (e *Employee) SetSalary(newSalary int) {
	e.salary = newSalary
}

// 4. Используйте методы с PointerReceiver в ситуациях когда
// 1) Изменения в работе метода над экземпляром должны быть видны на вызывающей стороне
// 2) Когда экземпляр достаточно "увесистый" то есть копировать его дорого с точки зрения расхода ресурсов
// 3) С ними может работать любой вид экземпляра

func main() {
	e := Employee{"Alex", 3000}
	fmt.Println("Before", e)
	e.SetName("Yan")
	fmt.Println("After", e)

	fmt.Println("Before", e)
	// 3. Вызов метода у ссылки на сотрудника
	e.SetSalary(5000)
	fmt.Println("After", e)
}
