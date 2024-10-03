package main

import "fmt"

// 1. Создадим тип Rectangle
type Rectangle struct {
	length, width int
}

// 2. создадим конструктор для Rectangle
// Для имен конструкторов в Go договаорились использовать функции со следующими названиями
// * New() если данный конструктор один на файл (в файле присутствует одна структура)
// * New<StructName()> если в файле присутствуют разные структуры

// 3. В Go принято возвращать из конструктора не сам экземпляр, а ссылку на него
func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

// 4. Добавим два метода
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := NewRectangle(10, 20)
	fmt.Printf("type is %T and value is %v\n", r, r)
	fmt.Println("Perimeter", r.Perimeter())
	fmt.Println("Area", r.Area())
}
