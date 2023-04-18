package main

import "fmt"

func main() {

	fmt.Println("For Biasa")
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka ", i)
	}
	fmt.Println("=====================")

	fmt.Println("For Dengan Argumen Hanya Kondisi")
	i := 1

	for i <= 5 {
		fmt.Println("Angka ", i)
		i++
	}
	fmt.Println("=====================")

	fmt.Println("Penggunaan Keyword break & continue")
	for i := 0; i <= 15; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println("Angka ", i)
	}
	fmt.Println("=====================")
}
