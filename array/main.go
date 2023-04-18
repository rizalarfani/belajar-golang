package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Rijal"
	names[1] = "Arfani"
	names[2] = "Cek"
	names[3] = "asdads"

	fmt.Println(names[0], names[1])

	var fruits = [4]string{"apple", "grace", "banana", "melo"}

	fmt.Println("Panjang", len(fruits))

	// Array Tanpa jumlah element
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Panjang", len(numbers))
	fmt.Println("Coba ", numbers[0])

	// Array Multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1[0][2])
	fmt.Println("numbers2", numbers2)

	for i := 0; i < len(fruits); i++ {
		fmt.Println("element ", i, fruits[i])
	}

	// Perulangan For Range, ini sama dengan foreach
	for i, fruit := range fruits {
		fmt.Println("element ", i, fruit)
	}
}
