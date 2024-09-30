package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var age int
	fmt.Println("My age:", age)

	var height int = 183
	fmt.Println("My height:", height)
	var uid = 12345

	fmt.Println("My uid:", uid)
	fmt.Printf("%T\n", uid)

	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 2, true

	fmt.Println("Total:", 2*quantity*(price+tax))

	var cost float32 = 275.00
	var total float32 = 27.50
	cost = 300
	fmt.Println(cost + total)

	var value = 275.00
	var taxes float32 = 27.50
	fmt.Println(value + float64(taxes))

	result := false
	value, new_value := 3.12, 121
	fmt.Println("Results:", result, new_value)

	var (
		number int
		s      string
	)

	fmt.Scan(&number)
	fmt.Scan(&s)
	fmt.Println(number, s)
}
