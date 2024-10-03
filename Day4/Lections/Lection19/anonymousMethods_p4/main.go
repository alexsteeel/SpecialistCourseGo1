package main

import (
	"fmt"
)

type University struct {
	city string
	name string
}

// 1. Данный метод связан только со структурой University
func (u *University) FullInfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s\n", u.name, u.city)
}

// 2. В структуру Professor встроены поля структуры University (методы тоже переходят)
type Professor struct {
	fullName string
	age      int
	University
}

func (p Professor) DetailInfo() {
	fmt.Println(p.fullName)
}

func main() {
	prof := Professor{
		fullName: "Stepan Bobrov",
		age:      45,
		University: University{
			city: "Moscow",
			name: "DMSTU",
		},
	}

	// 3. Вызываем метод структуры University через экземпляр профессора
	prof.FullInfoUniversity()

	// Вызываем свой метод
	prof.DetailInfo()
}
