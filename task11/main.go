package main

import (
	"fmt"
)

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {

	arr := []int{5, 2, 6, 3, 1, 4}

	fmt.Println(arr)

	reverse(arr)

	fmt.Println(arr)
}
