package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz!")
	fmt.Printf("Enter your name:\n")
	var name string
	fmt.Scan(&name) // user input
	fmt.Printf("Hello, %v! Welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	fmt.Println(age >= 10) // returns bool
}
