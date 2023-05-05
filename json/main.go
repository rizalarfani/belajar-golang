package main

import (
	"encoding/json"
	"fmt"
)

// Tag `json:"name"` digunakan untuk mapping informasi json ke property yang bersangkutan
type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("username :", data.FullName)
	fmt.Println("age :", data.Age)

	// Decode array json ke array object

	var jsonString2 = `[
		{"Name": "john wick", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data2 []User

	var err1 = json.Unmarshal([]byte(jsonString2), &data2)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}

	fmt.Println("user 1:", data2[0].FullName)
	fmt.Println("user 1:", data2[1].Age)

	// Encode object
	var object = []User{
		{
			"Rijal", 23,
		},
		{
			"Arfani", 23,
		},
	}
	var jsonData2, err2 = json.Marshal(object)

	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	var jsonString3 = string(jsonData2)
	fmt.Println(jsonString3)
}
