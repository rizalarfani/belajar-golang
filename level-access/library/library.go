package library

import "fmt"

func SayHello() {
	fmt.Println("hello")
	introduce("Rijal arfani")
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}
