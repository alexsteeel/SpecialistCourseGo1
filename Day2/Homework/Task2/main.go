package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Введите 5 чисел через пробел.")
	numbers := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numbers = scanner.Text()
	}

	array := strings.Fields(numbers)

	var asc, desc, constant bool
	for i := 0; i < 4; i++ {
		if array[i] > array[i+1] {
			desc = true
		} else if array[i] < array[i+1] {
			asc = true
		} else {
			constant = true
		}
	}

	switch {
	case asc && desc && constant:
		fmt.Println("RANDOM")
	case asc && constant:
		fmt.Println("WEAKLY ASCENDING")
	case asc:
		fmt.Println("ASCENDING")
	case desc && constant:
		fmt.Println("WEAKLY DESCENDING")
	case desc:
		fmt.Println("DESCENDING")
	case constant:
		fmt.Println("CONSTANT")
	default:
		fmt.Println("UNKNOWN")
	}
}
