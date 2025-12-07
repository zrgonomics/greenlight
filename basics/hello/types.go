package main

import "fmt"

func IntroTypes() {
	sigma := 'Σ'            // rune, alias for int32, holds a Unicode character
	pi := 3.141592653589793 // float64 by default
	fmt.Println("Sigma: %b, Pi: %f", sigma, pi)

	var a4 [4]int // array of 4 integers, each index initializes to 0
	a4[0] = 42
	// array literal, takes the size of the initialized values
	a5 := [...]int{1, 2, 3, 4, 5}

	fmt.Println("Array a4:", a4)
	fmt.Println("Array a5:", a5)

	// slice of integers, length 4, each index initializes to 0
	s4 := make([]int, 4)
	s4[0] = 24

	s5 := []int{5, 4, 3, 2, 1}     // slice literal
	s5 = append(s5, 0, -1, -2, -3) // append more values to the slice

	p, q := learnMemory()
	fmt.Println("Pointer p:", *p)
	fmt.Println("Pointer q:", *q)
}

func learnMemory() (p, q *int) {
	p = new(int)         // Allocates memory, returns pointer to zeroed int
	s := make([]int, 20) // Allocates slice of 20 ints, returns slice header
	s[3] = 7
	r := -2
	return &s[3], &r
}
