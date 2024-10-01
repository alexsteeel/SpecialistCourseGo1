package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Bits uint8

const (
	ASC Bits = 1 << iota
	WASC
	DESC
	WDESC
	CONST
	RAND
)

func Set(b, flag Bits) Bits   { return b | flag }
func Clear(b, flag Bits) Bits { return b &^ flag }
func Has(b, flag Bits) bool   { return b&flag != 0 }

func main() {

	fmt.Println("Введите 5 чисел через пробел.")
	numbers := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		numbers = scanner.Text()
	}

	array := strings.Split(numbers, " ")
	current := array[0]
	next := array[1]

	var b Bits
	b = Set(b, RAND)

	if current > next {
		b = Set(b, DESC)
		b = Set(b, WDESC)
	} else if current < next {
		b = Set(b, ASC)
		b = Set(b, WASC)
	} else {
		b = Set(b, WDESC)
		b = Set(b, WASC)
		b = Set(b, CONST)
	}

	for i := 2; i < 5; i++ {
		current = next
		next = array[i]

		if current > next {
			b = Clear(b, ASC)
			b = Clear(b, WASC)
			b = Clear(b, CONST)
		} else if current < next {
			b = Clear(b, DESC)
			b = Clear(b, WDESC)
			b = Clear(b, CONST)
		} else {
			b = Clear(b, ASC)
			b = Clear(b, DESC)
		}
	}

	for _, flag := range []Bits{CONST, ASC, WASC, DESC, WDESC, RAND} {
		if Has(b, flag) {
			switch flag {
			case ASC:
				fmt.Println("ASCENDING")
			case WASC:
				fmt.Println("WEAKLY ASCENDING")
			case DESC:
				fmt.Println("DESCENDING")
			case WDESC:
				fmt.Println("WEAKLY DESCENDING")
			case CONST:
				fmt.Println("CONSTANT")
			case RAND:
				fmt.Println("RANDOM")
			default:
				fmt.Println("UNKNOWN")
			}
			break
		}
	}
}
