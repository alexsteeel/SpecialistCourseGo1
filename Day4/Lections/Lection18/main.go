package main

import "fmt"

// 1. Вложенные структуры (вложение структур)
// Это использование одной структуры как тип поля для другой структуры
type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

// 4. Встроенные структуры (когда мы добавляем поля одной структуры к другой без имени)
type Professor struct {
	firstName  string
	lastName   string
	age        int
	greatWorks string
	University // в этом месте происходит добавление всех полей структуры University
}

func main() {
	std := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "good university",
			infoLong:  "It is very good university",
		},
	}
	fmt.Println(std)

	// 3. Получение доступа к вложенным полям структуры
	fmt.Println("firstName", std.firstName)
	fmt.Println("lastName", std.lastName)
	fmt.Println("yearBased of Uni", std.university.yearBased)

	// 5. Создание экземпляров структур с встроенной структурой
	prof := Professor{
		firstName:  "Anatoly",
		lastName:   "Smirnov",
		age:        35,
		greatWorks: "Ultimate GO programming",
		// papers map[string]string - добавление этого поля делает структуру несравнимой
		University: University{
			yearBased: 1734,
			infoShort: "old university",
			infoLong:  "very very old",
			age:       2024 - 1734,
		},
	}

	fmt.Println("firstName", prof.firstName)
	fmt.Println("yearBased", prof.yearBased)
	fmt.Println("another path to yearBased", prof.University.yearBased)

	fmt.Println("age of professor", prof.age)
	fmt.Println("age of university", prof.University.age)

	// 7. Сравнение экземпляров
	profLeft := Professor{}
	profRight := Professor{}
	fmt.Println(profLeft == profRight)

	// 8. Если хотя бы одно из полей структуры несравнимо, то и вся структура несравнима
}
