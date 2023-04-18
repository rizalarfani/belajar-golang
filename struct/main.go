package main

import "fmt"

func main() {
	// Deklarasi struct
	// Embedded Struct
	type student struct {
		person
		grade int
	}

	var struct1 student
	struct1.person.name = "Rijal Arfani"
	struct1.grade = 4.0
	struct1.person.age = 23

	fmt.Println("name  :", struct1.name)
	fmt.Println("grade :", struct1.grade)
	fmt.Println("age :", struct1.person.age)

	// Anonymous struct
	fmt.Println("Anonymous struct")
	var struct2 = struct {
		student
	}{}

	struct2.person.name = "Rijal arfani"
	fmt.Println("name  :", struct1.name)

	// Kombinasi Slice & Struct
	fmt.Println("Kombinasi Slice & Struct")
	var allPersons = []person{
		{
			name:    "Rijal Arfani",
			age:     23,
			address: "Adiwerna",
		},
		{
			name:    "Atik",
			age:     23,
			address: "Tarub",
		},
	}

	for _, person := range allPersons {
		fmt.Println(person.name, "age is ", person.age, "address is", person.address)
	}
}

type person struct {
	name    string
	age     int
	address string
}
