package main

import "fmt"

func main() {
	var number1 int = 5
	var number2 *int = &number1

	fmt.Println("Number1 (value) :", number1)
	fmt.Println("Number1 (Alamat) :", &number1)

	fmt.Println("Number2 (value) :", *number2)
	fmt.Println("Number2 (Alamat) :", number2)

	fmt.Println("Parameter Pointer")
	var number = 12
	fmt.Println("before :", number)
	change(&number, 50)
	fmt.Println("after :", number)
}

// Parameter Pointer
func change(original *int, value int) {
	*original = value
}
