package main

import "fmt"

func main() {
	helloWorld1 := "Hello, World!"
	fmt.Println(helloWorld1)
	var helloWorld2 = "Hello, world2!"
	fmt.Println(helloWorld2)
	const hw = "Hello, World __ hw!"
	fmt.Println(hw)
	var assignFromConst = hw
	fmt.Println(assignFromConst)
	fmt.Println("Hello, World!")  // This should print an output
}