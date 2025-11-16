package main

import (
	"fmt"
)

func add(numbers ...int) int {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(add(a, b))

}
