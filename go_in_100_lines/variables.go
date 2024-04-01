package main

import "fmt"

/* Declare a single variable */
var a int

/* Declare a group of variables */
var (
	b bool
	c float32
	d string
)

// Consts can also be declared like so
const HELLO_WORLD string = "Hello World!"

func main() {
	a = 42                  // Assign single value
	b, c = true, 32.0       // Assign multiple values
	d = "string"            // Strings must contain double quotes
	fmt.Println(a, b, c, d) // 42 true 32 string

	// Can also let go figure out the var type:

	x := "Hello"
	y := 21
	z := false

	fmt.Println(x, y, z)
	fmt.Println(HELLO_WORLD)
}
