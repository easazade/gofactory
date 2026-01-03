package main

import "fmt"
import "rsc.io/quote/v4"
import "examples.ali.com/greetings"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(quote.Opt())
	fmt.Println(greetings.Hello("Ali"))
}
