package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("request.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(content))

}
