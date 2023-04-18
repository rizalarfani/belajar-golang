package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var names = []string{"Rijal", " Arfani"}
	printMessage("Hello", names)

	fmt.Println("Fungsi Dengan Return Value / Nilai Balik")
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number:", randomValue)

	fmt.Println("Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi")
	divideNumber(4, 0)
}

// Function biasa
func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, "")
	fmt.Println(message, nameString)
}

// Fungsi Dengan Return Value / Nilai Balik
func randomWithRange(min, max int) int {
	value := randomizer.Int()%(max-min+1) + min
	return value
}

// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}
