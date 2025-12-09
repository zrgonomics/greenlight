package main // Declares an executable

import (
	"fmt"
	"strconv"
)

func BasicErrorHandling() {
	m := map[int]string{1: "one", 2: "two", 3: "three"}

	// ', ok' idiom is used to tell if something worked
	// 4 isn't in the map, so ok will be false
	if x, ok := m[4]; !ok {
		fmt.Println("Nobody's home")
	} else {
		fmt.Print(x)
	}

	if _, err := strconv.Atoi("non-int"); err != nil {
		fmt.Println(err)
	}
}
