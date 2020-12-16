package main

import (
	"encoding/json"
	"fmt"
)

// Book struck
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// struck instance
	book := Book{Title: "Title", Author: "Author"}
	fmt.Println("⬇️	Struct Version")
	fmt.Printf("%+v\n", book)
	// Marshal the instance to byte array
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("⬇️	JSON Version")
	fmt.Println(string(byteArray))
}
