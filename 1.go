package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Go 101")
	fmt.Println( quote.Hello() )
	fmt.Println( quote.Glass() )
	fmt.Println( quote.Go() )
	fmt.Println( quote.Opt() )
}