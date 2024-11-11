package main

import (
    "fmt"
    "stackup.dev/intro-to-golang/goFun"
)
func main() {
	sum := goFun.Add(5, 8)
	fmt.Println("The sum is", sum)

	greetings := goFun.Hello("Mwihoti");
	fmt.Println(greetings)
}