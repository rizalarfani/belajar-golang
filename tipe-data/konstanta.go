package main

import "fmt"

func main() {
	const firstName = "Rijal"
	lastName := "Arfani"

	fmt.Println("Nama Depan :", firstName)
	fmt.Println("Nama Belakang :", lastName)

	// Multi Constanta
	const (
		square       = "kotak"
		isToday bool = true
		numeric int8 = 1
	)
}
