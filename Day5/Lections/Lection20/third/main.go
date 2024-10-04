package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describe() {
	fmt.Printf("%s and years old %d\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	switch v := i.(type) { // Присваиваем переменной v значение, лежащие под предполагаемым интерфейсом
	case string:
		fmt.Println("This is stirng")
	case Describer:
		fmt.Println("This is describer")
		v.Describe()
	case Student:
		fmt.Println("This is student")
		v.GetName()
	}
}

func FindStudent(i interface{}) {
	switch v := i.(type) { // Присваиваем переменной v значение, лежащие под предполагаемым интерфейсом
	case string:
		fmt.Printf("This is string %v\n", v)
	case Student:
		fmt.Println("This is student")
		v.GetName()
	}
}

func (st Student) GetName() {
	fmt.Println("Name:", st.name)
}

func main() {
	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(std)
	TypeFinder(nil)

	FindStudent("golang")
	FindStudent(std)

	var describeStudent Describer = std
	FindStudent(describeStudent)
}
