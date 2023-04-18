package main

import (
	"fmt"
)

func main() {

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["maret"] = 100

	if chicken != nil {
		fmt.Println("Januari ", chicken["januari"])
		fmt.Println("Maret ", chicken["maret"])
	}

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    100,
		"april":    130,
		"mei":      150,
	}

	fmt.Println(chicken1, chicken2)

	// Membuat MAP Tanpa nilai awal
	var chicken3 = make(map[string]int)

	chicken3["maret"] = 100

	// Iterasi Item Map Menggunakan For - Range
	fmt.Println("Iterasi Item Map Menggunakan For - Range")

	var product = map[string]string{
		"name":        "Mouse",
		"description": "Ini adalah mouse untuk komputer saya",
		"stock":       "10",
	}

	for key, value := range product {
		fmt.Println(key, " \t", value)
	}

	// Deteksi Keberadaan Item Dengan Key Tertentu
	fmt.Println("Deteksi Keberadaan Item Dengan Key Tertentu")
	var value, isExistChicken = chicken2["maret"]

	if isExistChicken {
		fmt.Println("Tersedia ", value)
	} else {
		fmt.Println("Item not exists")
	}

	// Kombinasi Slice & Map
	fmt.Println("Kombinasi Slice & Map")

	var products = []map[string]string{
		{
			"name":        "Komputer",
			"description": "ini adalah komputer saya",
			"qty":         "1",
		},
		{
			"name":        "Monitor",
			"description": "Monitor untuk komputer saya",
			"qty":         "2",
		},
		{
			"name":        "Keybord",
			"description": "Keybord untuk komputer saya",
			"qty":         "1",
		},
	}

	if products != nil {
		fmt.Println("Amount :", len(products))
		for _, product := range products {
			fmt.Println(product["name"], " \t", product["qty"])
		}
	}
}
