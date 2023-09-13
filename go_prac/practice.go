package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter yo name: ")

	var name string

	fmt.Scan(&name) //storing in memory referencing the variable

	fmt.Printf("Hello, %v, welcome to the game!", name)

	fmt.Printf("\nEnter your age: ")
	var age uint

	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("yayy u can playy")
	} else {
		fmt.Println("oh no u can't play")
		return
	}

	fmt.Printf("What is better, tea or coffee?")
	var answer string
	fmt.Scan(&answer)

	fmt.Println(answer)

	if answer == "coffee" {
		fmt.Println("yayy")
	} else if answer == "COFFEE" {
		fmt.Println("YAYY")
	} else {
		fmt.Println("yoo")
	}

}
