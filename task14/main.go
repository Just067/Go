package main

import "fmt"

// BubbleSort (сортировка пузырьком по убыванию)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		flag := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

// ShakerSort (шейкерная сортировка)
func ShakerSort(arr []int) {
	left, right := 0, len(arr)-1
	flag := 1

	for left < right && flag > 0 {
		flag = 0
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = 1
			}
		}
		right--
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				flag = 1
			}
		}
		left++
		if flag == 0 {
			break
		}
	}
}

// InsertionSort (сортировка вставками)
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}

// SelectionSort (сортировка выбором)
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		index := 0
		for j := 1; j < n-i; j++ {
			if arr[j] > arr[index] {
				index = j
			}
		}
		arr[index], arr[n-i-1] = arr[n-i-1], arr[index]
	}
}

// ShellSort (сортировка Шелла)
func ShellSort(arr []int) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = temp
		}
	}
}

// GnomeSort (гномья сортировка)
func GnomeSort(arr []int) {
	n := len(arr)
	i, j := 1, 2
	for i < n {
		if arr[i-1] < arr[i] {
			i = j
			j++
		} else {
			arr[i-1], arr[i] = arr[i], arr[i-1]
			i--
			if i == 0 {
				i = j
				j++
			}
		}
	}
}

// Функция для вывода массива
func printArray(arr []int, name string) {
	fmt.Printf("%s: %v\n", name, arr)
}

// Функция для копирования массива
func copyArray(original []int) []int {
	copyArr := make([]int, len(original))
	copy(copyArr, original)
	return copyArr
}

func main() {
	arr := make([]int, 5)
	for i := 0; i < 5; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return
		}
	}

	fmt.Println("\nИсходный массив:", arr)
	fmt.Println("\nДемонстрация различных алгоритмов сортировки:")

	// Демонстрация каждой сортировки на копии исходного массива
	arrCopy := copyArray(arr)
	BubbleSort(arrCopy)
	printArray(arrCopy, "BubbleSort(desc)")

	arrCopy = copyArray(arr)
	ShakerSort(arrCopy)
	printArray(arrCopy, "ShakerSort")

	arrCopy = copyArray(arr)
	InsertionSort(arrCopy)
	printArray(arrCopy, "InsertionSort")

	arrCopy = copyArray(arr)
	SelectionSort(arrCopy)
	printArray(arrCopy, "SelectionSort")

	arrCopy = copyArray(arr)
	ShellSort(arrCopy)
	printArray(arrCopy, "ShellSort")

	arrCopy = copyArray(arr)
	GnomeSort(arrCopy)
	printArray(arrCopy, "GnomeSort")

}
