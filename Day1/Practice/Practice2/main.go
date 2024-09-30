package main

import "fmt"

const minusAge uint8 = 1

func main() {
	var name string
	var age uint8

	fmt.Println("Enter your name pls")
	fmt.Scan(&name)

	fmt.Println("Enter your age")
	fmt.Scan(&age)
	fmt.Printf("Your name is %s\n", name)
	fmt.Printf("Your age minus 1 is %d\n", age-minusAge)
}
