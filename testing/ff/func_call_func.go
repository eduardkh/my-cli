package main

import "fmt"

func one() string {
	return fmt.Sprintln("one run")
}

func two() string {
	return fmt.Sprintln("two run")
}

func three(one func() string, two func() string) string {
	return fmt.Sprintln("three run", one(), two())
}

func main() {
	fmt.Print(one())
	fmt.Print(two())
	fmt.Print(three(one, two))

}
