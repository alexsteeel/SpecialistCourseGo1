package main

import "fmt"

func main() {
	aaa := 1111

	var result = 0
	for i := 0; i < 100000000; i++ {
		ccc := aaa
		if result%2 == 0 {
			result += ccc
		} else {
			result -= ccc
		}
	}
	fmt.Println(result)
}
