package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Описание интерфейса
type BigWorld interface {
	IsBig() bool
}

// 2. наш тип-претендент, который должен иметь метод IsBig()
type MySuperString string

// 3. Реализация интерфейса
func (mss MySuperString) IsBig() bool {
	return utf8.RuneCountInString(string(mss)) > 10
}

func main() {
	sample := MySuperString("hellpjhtgufytd")
	var interfaceSample BigWorld = sample
	fmt.Println(interfaceSample.IsBig())

	// 4. Попытка присвоить значение с типажом, не удовлетворяющему интерфейсу, приводит к ошибке

}
