package main // Declares an executable

import (
	"fmt" 				// Standard library fmt
	"io"              	// Implements some I/O utility functions.
	m "math"          	// Math library, aliased to m
	"net/http"        	// Web server
	_ "net/http/pprof" 	// Profiling library imported only for side fx
	"os"               	// Basic I/O functions
	"strconv"  			// String conversions
)

// main() is special in Go, it's the entrypoint to an executable
// No declaring other functions within main
func main() {
	fmt.Println("Greenlight!")

	StartServer()
	greenlightUltra()
	IntroTypes()

	// Weirdly, Go has goto labels
	goto jump

	jump:
}

func greenlightUltra() {
	//var s string   // initializes to "", Go doesn't do null for primitives
	var i int // Initialize to 0... 32 or 64 bit depending on architecture
	var i32 int32  // Explicit 32-bit int.. int8, int16, int32, int64 work
	var u32 uint32 // Explicit unsigned 32-bit int
	var f float64  // 0.0...
	var b bool     // false...

	// Unused variables cause compile error, so use blank identifier, discards them
	_, _, _, _, _ = i32, u32, f, b, m.Pi // prevent "declared and not used" error

	//var p *int      // nil pointer
	//var slice []int // nil slices
	//var chunk byte // alias for uint8

	// Go has NO IMPLICIT TRUTHINESS. More verbose, less ambiguous.
	if i != 0 {
		fmt.Println("i initialized or not zero.")
	}

	result, err := intInitialized()

	if err != nil {
		fmt.Println(result)
	}
}

func intInitialized(nums ...int) (bool, error) {
	if nums != nil {}
	return true, fmt.Errorf("not an int")
}

// Go can have named return values, kinda neat
func namedReturns(x, y int) (z int) {
	z = x + y
	return // naked return, returns z cause we named it
}

func controlFlow() {
	// No parentheses around conditions
	if true {
		fmt.Println("Indubitably")
	} else {
		fmt.Println("Nope")
	}

	x := 42.0

	switch x {
	case 0:
	case 1, 2:
	case 42:
		fmt.Println("The answer to life, universe and everything")
	default: // Optional
		fmt.Println("Something else")
	}

	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println("data is a string:", c)
	case int:
		fmt.Println("data is an int:", c)
	default:
		fmt.Println("data is another type")
	}

	// Go's only looping construct
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	var numbers map[string]int {"one": 1, "two": 2}

	for k, v := range map[string]int{"one": 1, "two": 2} {
		fmt.Println("Key:", k, "Value:", v)
	}

	for _, v := 
}