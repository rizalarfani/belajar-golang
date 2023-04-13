package main

import "fmt"

func main() {

	var firstName, lastName, age string
	firstName, lastName, age = "Rijal", "Arfani", "23"

	fmt.Print("Nama : ", firstName, " ", lastName)
	fmt.Println()
	fmt.Print("Age : ", age)
	fmt.Println()

	address, phone := "Tembok Banjaran", "086426755210"

	fmt.Println("Alamat :", address)
	fmt.Println("Phone :", phone)
}
