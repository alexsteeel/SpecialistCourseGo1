package main

import (
	"fmt"
)

func main() {
	// 1. Указатели - переменная, хранящая в качестве значения адрес в памяти другой переменной
	// 2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable // &... - операция взятия адреса в памяти
	// выше создан указатель на переменную variable
	// В pointer лежит какой-то адрес for instance 0xc00008e040 - это место в памяти, где хранится занчение int32

	fmt.Printf("Type of pointer: %T\n", pointer)
	fmt.Printf("Value of pointer: %v\n", pointer)

	// 3. А какое нулевое значение для указателя?
	var zeroPointer *int //zeroPointer имеет значение nil (указатель в никуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)

	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	}

	// 4. Разыменовывание указателя (получение значения по адресу):
	// *pointer - возвращает значение, которое хранится по этому адресу
	var numericValue int = 32
	pointerToNumeric := &numericValue
	fmt.Printf("Value of pointerToNumeric: %v and it's address %v\n",
		*pointerToNumeric, pointerToNumeric)

	// 5. Создание указателей на нулевые значения типов
	var zeroVar int
	var zeroPointer2 *int = &zeroVar
	fmt.Printf("zero pointer %v and it's value %v\n", zeroPointer2, *zeroPointer2)

	zeroNewPoint := new(int) // создает под капатом zeroValue для int и возвращает адрес, где этот 0 хранится
	fmt.Printf("zero new pointer %v and it's value %v\n", zeroNewPoint, *zeroNewPoint)

	// 6. Изменение значения
	zeroPtoInt := new(int)
	fmt.Printf("Value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)

	*zeroPtoInt += 40
	fmt.Printf("New value of zeroPtoInt: %v and it's address %v\n", *zeroPtoInt, zeroPtoInt)

	b := 345
	a := &b
	c := &b
	*a++
	*c += 100
	fmt.Println("a:", *a, "b:", b, "c:", *c)

	// 7. Указательная арифметика ПОЛНОСТЬЮ ОТСУТСТВУЕТ
	// У вас на руках адрес одной ячейки - вы можете через этот адрес перейти в другие ячейки

	// 8. Передача указателей в функции
	// Большой прирост производительности за счет того, что передается не само значение, которое может быть очень большим
	// и которое по умолчанию копируется, а передается только адрес в памяти, за которым уже хранится какое-то значение

	sample := 1
	fmt.Println("Origin value of sample:", sample)
	changeParam(&sample)
	fmt.Println("Updated value of sample:", sample)

	// 9. Возврат указателя из функции
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("ptr1 type %T address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("ptr2 type %T address %v and value %v\n", ptr2, ptr2, *ptr2)

	arr := [3]int{1, 2, 3}
	fmt.Println("Before calling mutation", arr)
	mutation(&arr)
	fmt.Println("After calling mutation", arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println("Before calling mutation", arr2)
	mutationSlice(arr2[:])
	fmt.Println("After calling mutation", arr2)
}

// 8.1 Определение функции, принимающей указатель как параметр
func changeParam(value *int) {
	*value += 100
}

func returnPointer() *int {
	var numeric int = 321
	return &numeric // в момент возврата go перемещает данную переменную в кучу
}

// 10. Указатели на массивы
func mutation(arr *[3]int) {
	// (*arr)[1] = 900
	// (*arr)[2] = 1000
	// Но можно писать и так, потому что go сам разыменуют указатель на массив
	// из-за того что функция принимает *arr
	arr[1] = 900
	arr[2] = 1000
}

// 11. Лучше используйте слайсы (это идиоматично с точки зрения go)
func mutationSlice(sls []int) {
	sls[1] += 900
	sls[2] += 1000
}
