package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Введите число: ")
	fmt.Scan(&N)

	if N <= 0 {
		return
	}

	num1 := 0
	num2 := 1

	fmt.Print(num1, num2)
	for i := 0; i < N-2; i++ {
		res := num2 + num1
		fmt.Print(" ", res)
		num1 = num2
		num2 = res
	}

}
