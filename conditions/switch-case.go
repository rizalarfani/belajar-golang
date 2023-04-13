package main

import "fmt"

func main() {
	point := 6

	switch point {
	case 8:
		fmt.Printf("Perfect")
	case 7:
		fmt.Println("Awasome")
	default:
		fmt.Println("Not Bad")
	}
}
