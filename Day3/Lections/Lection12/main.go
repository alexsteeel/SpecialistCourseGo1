package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var name string

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		// Команда захвата потока ввода и сохранения в буфер
		// (захват идет до символа новой строки)
		name = input.Text()
		// Метод Text() возвращает элементы, помещенные в буфер в виде строки
	}
	fmt.Println(name)

	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}

	// Преобразование строкового литерала к чему-нибудь числовому
	numStr := "10"
	numInt, _ := strconv.Atoi(numStr)
	fmt.Printf("Number: %d and type is %T\n", numInt, numInt)

	numInt64, _ := strconv.ParseInt(numStr, 10, 64)
	fmt.Printf("Number: %d and type is %T\n", numInt64, numInt64)

	numFloat32, _ := strconv.ParseFloat(numStr, 32)
	// Но это 64 разрядное число будет без проблем гарантированно приводиться к float32
	fmt.Printf("Number: %.3f and type is %T\n", numFloat32, float32(numFloat32))

	// Пример bufio.NewReader
	fmt.Println("bufio.NewReader example")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Fields(line)
	fmt.Printf("data value: %v and as is: %#v\n", data, data)
}
