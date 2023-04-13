package main

import "fmt"

func main() {
	name := new(string)
	*name = "Rijal Arfani"

	fmt.Println("Value :", *name)
	fmt.Println("Alamat memori :", name)
}
