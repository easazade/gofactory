package main

import (
	"fmt"

	"examples.ali.com/greetings"
	"rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Opt())
	fmt.Println(greetings.Hello("Ali"))
}
