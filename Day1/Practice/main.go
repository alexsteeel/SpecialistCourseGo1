package main

import "fmt"

func main() {
	var name string
	var age uint8

	fmt.Println("Enter your name pls")
	fmt.Scan(&name)

	fmt.Println("Enter your age")
	fmt.Scan(&age)
	fmt.Println("Your name is", name)
	fmt.Println("Your age minus 1 is", age-1)
}
