package main

import "fmt"

/*
1. Явная функция - локально опеределенный блок кода, имеющий имя (явное определение)
Функцию необходимо определить + функцию необходимо вызвать
2. Сигнатура функций и их определениеЭ
func functionName(params type) typeReturnValue {
	function body
}

3. Показать документацию пакета
go doc
*/

// Пакет
func main() {
	//myFunction()

	fmt.Println("call function")
	res := add(10, 20)
	fmt.Println(res)

	per, square := rectangleParameters(12, 4)
	fmt.Println(per, square)

	perNamed, squareNamed := namedReturn(12, 6)
	fmt.Println(perNamed, squareNamed)

	helloVariadic(1, 2, 3)

	someStrings(1, 2, "123", "456")

	fmt.Println(sumVariadic(1, 2, 3))
}

// 3. Простейшая функция (определить функцию можно как до момента ее вызова в функции main),
// так и в любом месте пакета главное, чтобы она была определена в принципе где-нибудь

// add function return of two numbers
func add(a int, b int) int {
	result := a + b
	return result
}

// 5. Функция с однотипными параметрами
func mult(a, b, c, d int) int {
	return a * b * c * d
}

// 6. Возврат больше чем одного значения (returnType1, returnType2 etc)
func rectangleParameters(a, b int) (int32, int) {
	perimeter := int32(2 * (a + b))
	area := a * b
	return perimeter, area
}

// 7. Именованный возврат значений
func namedReturn(width, height int) (perimeter int32, area int) {
	perimeter = int32(2 * (width + height))
	area = width * height
	return // не нужно указывать возвращаемые значения
}

// 8. При вызове оператора return функция прекращает свое выполнение и что-то возвращает
func funcWithReturns(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}

	if a == b {
		return a, true
	}

	return 0, false
}

// 9. Что вернет функция в случае если return не указан или он пустой
func emptyReturn(a int) {
	fmt.Println("I am empty return with parameter a", a)
}

// 10. Variadic parameters (континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Printf("value %v and it's type %T\n", a, a)
}

// 11. Смешивание параметров с variadic (
//  1. Континуальный параметр всегда самый последний
//  2. variadic parameter - один на всю функцию
//
// )
func someStrings(a, b int, words ...string) {
	fmt.Println(a+b, words)
}

// 12. Передача слайса или variadic parameters
func sumVariadic(nums ...int) int {
	var sumOfInputs int
	for _, number := range nums {
		sumOfInputs += number
	}
	return sumOfInputs
}

func sumSlice(nums []int) int {
	var sumOfInputs int
	for _, number := range nums {
		sumOfInputs += number
	}
	return sumOfInputs
}
