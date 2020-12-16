package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Session struct
type Session struct {
	Authenticated bool   `json:"authenticated"`
	Token         string `json:"token"`
}

func main() {
	// get json file content
	jsonString, err := ioutil.ReadFile("request.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(jsonString))
	// Session struct instance
	var reading Session
	// Unmarshal from byte slice to Session struct instance
	err = json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
}
