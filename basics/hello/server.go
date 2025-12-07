package main

import "fmt"

type Server struct {
	Port int8
}

var internalCache = make(map[string]string)

func StartServer() { // TitleCase name = public, exported
	fmt.Println("Starting server...")
	helper()
}

func helper() { // lowercase name = private to the package, not exported
	fmt.Println("Private helper function ran")
}
