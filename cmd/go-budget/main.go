package main

import (
	"fmt"
	"go-budget/pkg/io"
)

type Data struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func main() {
	data := Data{Field1: "Hello, JSON!", Field2: 42}

	err := io.SaveJSONToFile(data, "data.json")
	if err != nil {
		fmt.Println("Error saving JSON to file: ", err)
		return
	}

	fmt.Println("Data saved to data.json successfully.")
}
