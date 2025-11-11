package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	
	if b == 0 {
		fmt.Println(a+b)
		fmt.Println(a-b)
		fmt.Println(a*b)
		fmt.Println("Деление на ноль невозможно!")
	} else {
		fmt.Println(a+b)
		fmt.Println(a-b)
		fmt.Println(a*b)
		fmt.Printf("%.4f\n", float64(a)/float64(b))
		fmt.Println(a % b)
	}
}
