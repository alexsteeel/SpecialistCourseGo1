package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)

	//  Тип данных можно получить через %T форматирование
	fmt.Printf("type of a: %T and type of b: %T\n", a, b)

	//  Узнаем, сколько байт в памяти занимает тип
	// При использовании короткого объявления int - платформо-зависимый
	fmt.Printf("type of a %T and size in memory in bytes: %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("type of b %T and size in memory in bytes: %d\n", b, unsafe.Sizeof(b))

	// Если выполняются арифметические операции над int и intX
	// то обязательно нужно использовать механизм приведения типов,
	// так как int != int64
	var c64 int64 = 16_213_897
	var usualInt int = 123_567
	fmt.Println(c64 + int64(usualInt))

	// Аналогичным образом устроены uint8, uint16, uint32, uint64

	// Numerics. Float
	// float32, float64
	floatFirst, floatSecond := 5.67, 12.54

	fmt.Printf("type of floatFirst %T and size in memory in bytes: %d\n", floatFirst, unsafe.Sizeof(floatFirst))
	sum := floatFirst + floatSecond
	sub := floatFirst - floatSecond
	fmt.Printf("Sum: %.3f, Sub: %.3f\n", sum, sub)

	fmt.Println(math.Pow(floatFirst, floatSecond))

	// Numerics. Complex
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	// Strings.  Строка - это набор байт
	name := "Федя"
	lastname := "Pupkin"
	fi := name + " " + lastname
	fmt.Println("Full name:", fi)

	// функция len() - кол-во байт в строке
	fmt.Println("Length of string", len(name))
	fmt.Println("Length of string", len(lastname))

	// Rune - руна. Это один utf-ный символ
	fmt.Println("Кол-во символов в строке:", utf8.RuneCountInString(name))
	fmt.Println("Кол-во символов в строке:", utf8.RuneCountInString(lastname))

	// Поиск подстроки в строке
	totalString, subString := "ABCDEFG", "CDE"
	fmt.Println(strings.Contains(totalString, subString))

	// rune - это alias int32
	var sampleRune rune
	var anotherRune = 'Q' // Для инициализации руны использовать одинарные кавычки
	var thirdRune = 234
	fmt.Println(sampleRune)
	fmt.Printf("Rune as char %c\n", sampleRune)
	fmt.Printf("Rune as char %c\n", anotherRune)
	fmt.Printf("Rune as char %c\n", thirdRune)

	// -1 if first < second, 0 if first == second, 1 if first > second
	fmt.Println(strings.Compare("abcd", string('a')))

	var aByte byte // alias uint8
	fmt.Println("Byte", aByte)
}
