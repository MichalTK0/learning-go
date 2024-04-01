package main

import (
	"fmt" 
	"strings"
)

func main() {

	// Make a slice of strings
	slc := []string{
		"test1", "test2", "test3",
	}

	// For loops - pretty standard
	for _, v := range slc {
		fmt.Printf("%s, ", v)
	}
	fmt.Println("\n")

	// The print above ends up with an extra comma.. luckily go has something similar to python's join
	result := strings.Join(slc, ", ")
	fmt.Println(result)

	// While loops - they don't exist. You just have to use a for loop in this way:
	fmt.Println("\n")
	count := 0

	for count < len(slc) {
		fmt.Printf("%s ", slc[count])
		count += 1
	}

}
