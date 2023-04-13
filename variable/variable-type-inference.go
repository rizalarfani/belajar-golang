package main

import "fmt"

var name string = "asd"
var age = 10

func main() {
	// Tanpa var, tanpa tipe data, menggunakan perantara ":="
	// Tetapi deklarasi dengan ini hanya bisa dilakukan di blok fungsi
	firstName := "Rijal"
	lastName := "Arfani"

	fmt.Printf("Hello %s %s \n", firstName, lastName)
}
