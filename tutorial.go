package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz!")
	var name string
	fmt.Scan(&name) // user input
	fmt.Printf("Your name is %v!", name)
}
