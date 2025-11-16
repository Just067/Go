package main

import (
	"errors"
	"fmt"
)

func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	quotient := a / b
	remainder := a % b
	return quotient, remainder, nil
}

func main() {
	var a, b int

	fmt.Scan(&a, &b)
	q, r, err := divmod(a, b)

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(q, r)
}
