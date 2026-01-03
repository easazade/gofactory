package main

import "fmt"

/*
Package-level (global) variables
*/

// Explicit type + value
var appName string = "DemoApp"

// Type inferred
var version = "1.0.0"

// Multiple globals
var (
	debug   = true
	retries int
)

// Constant (compile-time)
const maxUsers = 100

/*
init runs automatically before main
You can have multiple init() functions across files
*/
func init() {
	retries = 3 // assigning to a global
	fmt.Println("init(): setup complete")
}

func main() {

	/*
	   Function-level variables
	*/

	// Short declaration (declare + assign)
	count := 10

	// Explicit type
	var name string = "Ali"

	// Declare first, assign later
	var score int
	score = 42

	// Multiple assignment
	a, b := 1, 2

	fmt.Println(appName, version, debug)
	fmt.Println(count, name, score, a, b)
}
