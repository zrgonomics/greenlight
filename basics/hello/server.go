package main

import "fmt"

func StartServer() { // TitleCase name = public, exported
	fmt.Println("Starting server...")
	helper()
}

func helper() { // lowercase name = private to the package, not exported
	fmt.Println("Private helper function ran")
}
