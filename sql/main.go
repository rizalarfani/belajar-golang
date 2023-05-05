package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	nim   string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select nim,name,grade from tbl_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var results []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.nim, &each.name, &each.grade)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		results = append(results, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range results {
		fmt.Println(each.nim)
	}
}

func main() {
	sqlQuery()
}
