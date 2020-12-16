package main

import (
	"encoding/json"
	"fmt"
)

// Book struck
type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

// Author struck
type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	// Author struck instance
	author := Author{Name: "Name", Age: 36, Developer: true}
	// Book struck instance
	book := Book{Title: "Title", Author: author}
	fmt.Println("⬇️	Struct Version")
	fmt.Printf("%+v\n", book)
	// Marshal the instance to byte array with indentation
	byteArray, err := json.MarshalIndent(book, "", "    ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("⬇️	JSON Version")
	fmt.Println(string(byteArray))
}
