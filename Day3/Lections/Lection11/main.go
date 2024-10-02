package main

import (
	"fmt"
)

func main() {
	var array [3]int
	array[0] = 100
	array[1] = 200
	array[2] = 300

	fmt.Println(array)
	const a = 6
	fmt.Println(a)
	var array1 [a]int
	fmt.Println(array1)

	// Слайсы (они же срезы)
	// Слайсы - это динамическая обвязка над массивом
	startArr := [4]int{10, 20, 30, 40}
	var startSlice []int = startArr[0:2] // Слайс инициализируется пустыми квадратными скобками
	fmt.Println("Slice[0:2]", startSlice)
	// Создали слайс на основе уже существующего
	startArr[0] = 100
	fmt.Println("Slice[0:2]", startSlice)

	// 2. Создание слайса без явной инициализации массива
	secondSlice := []int{15, 20, 25, 30, 70}
	fmt.Println("second slice", secondSlice)

	// 3. Изменение элементов среза
	originArr := [...]int{30, 40, 50, 60, 70, 80}
	thirdSlice := originArr[1:4] // набор ссылок на элементы нижележащего массива
	for i, _ := range thirdSlice {
		thirdSlice[i]++
	}

	fmt.Println("origin:", originArr)
	fmt.Println("slice:", thirdSlice)

	// 4. Один массив и два производных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]
	fmt.Println("Before changing. Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)
	fSlice[3]++
	sSlice[1]++
	fmt.Println("After changing. Arr:", originArr, "fSlice:", fSlice, "sSlice:", sSlice)

	// 5. Срез как встроенный тип
	// type slice struct {
	//	Length int
	//  Capacity int
	//  ZeroElement *byte
	// }

	// 6. Длина и емкость слайса
	wordsSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", wordsSlice, "Length:", len(wordsSlice), "Capacity:", cap(wordsSlice))

	// Добавить новый элемент
	wordsSlice = append(wordsSlice, "four")
	fmt.Println("slice:", wordsSlice, "Length:", len(wordsSlice), "Capacity:", cap(wordsSlice))

	// Capacity (cap) или емкость слайса - это значение, показывающее сколько элементов можно
	// хранить в слайсе, не выделяя дополнительной памяти
	// Если нет места под новый элемент, то выделяется память, равная n*2
	// где n - это размер емкости до изменения
	// cap = 3: 3 -> 6 -> 12 ...
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	// Важно: после выделения памяти под новый массив ссылки со старого будут перенесены в новый
	// Пример
	numArr := []int{1, 2}
	numSlice := numArr[:]
	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 0
	fmt.Println(numArr)
	fmt.Println(numSlice)

	// 7. Как создавать слайсы наиболее эффективно
	// make() - это функция позволяющая более детально создавать срезы
	sl := make([]int, 10, 15)

	// []int - тип коллекции
	// 10 - длина
	// 15 - емкость
	// Сначала инициализируется arr = [15]int
	// Затем по нему делается срез от 0 до 10 arr[0:10]
	// После этого он возвращается
	fmt.Println(sl)

	// 8. Добавление элементов в срез
	myWords := []string{"one", "two", "three"}
	fmt.Println("my words:", myWords)
	anotherWords := []string{"four", "five", "six"}
	myWords = append(myWords, anotherWords...)
	myWords = append(myWords, "seven", "eight")
	fmt.Println("After change:", myWords)

	// 9. Многомерный срез
	mySlice := [][]int{
		{1, 2},
		{2, 3, 4, 5},
		{10, 20, 30},
		{},
	}

	fmt.Printf("Values: %v as is %#v\n", mySlice, mySlice)

	// 10. Слайсы рун
	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	myStrFromRunes := string(runeHexSlice)
	fmt.Println("Runes:", runeHexSlice, "myStr", myStrFromRunes)

	// 10.1 Руны как литералы
	runeLiteralSlice := []rune{'V', 'a', 's', 't'}
	myStrFromLiteralRunes := string(runeLiteralSlice)
	fmt.Println("Runes:", runeLiteralSlice, "myStr", myStrFromLiteralRunes)

	// 10.2 Вычисление емкости строки бессмысленно, так как строка - базовый тип
	fmt.Println(cap([]rune("Vasya")))

	// 10.3 Строки неизменяемые (immutable), а слайсы рун изменяемые (mutable)
	firstName := "John"
	fmt.Println(firstName)

	firstSliceName := []rune("John")
	firstSliceName[0] = 'B'
	fmt.Println(string(firstSliceName))

	// 10.4 Синтаксический сахар (можно использовать десятичное представление байтов)
	myDecimalByteSlice := []byte{100, 101, 102, 103}
	myDecimalString := string(myDecimalByteSlice)
	fmt.Println("myDecimalString:", myDecimalString)
}
