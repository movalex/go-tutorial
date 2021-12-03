package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz!")
	fmt.Printf("Enter your name:\n")
	var name string
	fmt.Scan(&name) // user input
	fmt.Printf("Hello, %v! Welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint // unsigned integer
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("Yay, you can play game!\n")
	} else {
		fmt.Println("You can't play!")
	}
}
