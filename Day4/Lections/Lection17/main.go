package main

import "fmt"

// 1. Структура - заименованный набор полей (состояний) определяющий новый тип данных
// 2. Определение структуры (явное)
type Student struct {
	firstName string
	lastName  string
	age       int
}

func PrintStudent(std Student) {
	fmt.Println("====================")
	fmt.Println("firstName:", std.firstName)
	fmt.Println("lastName:", std.lastName)
	fmt.Println("age:", std.age)
}

// 11. Структура с анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

// 3. Если имеется ряд состояний одинакового типа, то можно сделать так;
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, courseNumber              int
}

func main() {
	// 4. Создать представителя (экземпляр) структуры
	student := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		age:       19,
	}
	fmt.Println(student)
	PrintStudent(student)

	student2 := Student{"Vasya", "Sidorov", 21} // порядок указания свойств такой же как и в структуре
	PrintStudent(student2)

	// 5. Что если не все поля структуры определены?
	student3 := Student{
		firstName: "Petya",
	}
	PrintStudent(student3)

	// 6. Анонимные структуры (структуры без имени)
	anonStudent := struct {
		age           int
		groupID       int
		professorName string
	}{
		age:           22,
		groupID:       12,
		professorName: "Stepan Stepanovich",
	}
	fmt.Println("anonStudent:", anonStudent)

	// 7. Доступ к полям (состояниям) и их модификация
	studMark := Student{"Mark", "Ivanov", 20}
	PrintStudent(studMark)
	studMark.age += 2
	PrintStudent(studMark)

	// 8. Инициализация пустой структуры
	emptyStudent1 := Student{}
	var emptyStudent2 Student

	PrintStudent(emptyStudent1)
	PrintStudent(emptyStudent2)

	// 9. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       21,
	}
	fmt.Println("Value of pointer:", studPointer)

	secondPointer := studPointer
	(*secondPointer).age += 10
	fmt.Println("Value of pointer after changing:", studPointer)

	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	// 10. Работа с доступом к полям через указатель
	fmt.Println("Age via (*...).age", (*studPointer).age)
	// неявно происходит разыменование и обращению к соответствующему полю
	fmt.Println("Age via .age", studPointer.age)

	// 11. Создание экземпляра с анонимными полями
	human := &Human {
		firstName: "Bob",
		lastName: "Smirnoff",
		string: "Add info",
		int: -1,
		bool: true,
	}

	fmt.Println(human)
	fmt.Println("Anon field string", human.string)
}
