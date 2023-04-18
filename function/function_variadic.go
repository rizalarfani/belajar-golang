package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3, 10, 100, 200)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	total := 0
	for _, val := range numbers {
		total += val
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
