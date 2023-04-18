package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"wick", "jason", "ethan"}

	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLenght5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	// filter ada huruf "o" : [jason]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)

	// filter jumlah huruf "5" : [jason ethan]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5)

}

// Alias Skema Closure
type filterCallback func(string) bool

func filter(data []string, callback filterCallback) []string {
	var result []string
	for _, val := range data {
		if filltered := callback(val); filltered {
			result = append(result, val)
		}
	}
	return result
}
