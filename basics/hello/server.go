package main

import "fmt"

type Server struct {
	Port int8
}

var internalCache = make(map[string]string)

// TitleCase name = public, exported
func StartServer() {
	fmt.Println("Starting server...")
	helper()
}

// lowercase name = private to the package, not exported
func helper() {
	fmt.Println("Private helper function ran")
}
