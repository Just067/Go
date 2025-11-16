package main

import "fmt"

func Stats(arr []int) (sum int, maxIdx int, maxVal int, positives []int, avg float64) {

	for i, val := range arr {
		sum += val

		if val > maxVal {
			maxVal = val
			maxIdx = i
		}

		if val > 0 {
			positives = append(positives, val)
		}
	}

	avg = float64(sum) / float64(len(arr))

	return
}

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}

	sum, maxIdx, maxVal, positives, avg := Stats(a[:])

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Индекс максимального элемента: %d\n", maxIdx)
	fmt.Printf("Максимальный элемент: %d\n", maxVal)
	fmt.Printf("Положительные элементы: %v\n", positives)
	fmt.Printf("Среднее арифметическое: %.2f\n", avg)

}
