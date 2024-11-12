package main

import (
    "fmt"
    "stackup.dev/intro-to-golang/goFun"
)

func is_legal_age(age int) (msg string, legal bool) {
	var m string
	var is_legal bool
	if age >= 18 {
		m = "You are now of legal age!"
		is_legal = true
		return m, is_legal
	}
	m = "You are a minor!"
	is_legal = false
	return m, is_legal
}

func is_rich(rich bool) string {
	var m string
	if rich == true {
		m = "You are rich motherfucker"
		return m
	}
	m = "So pal you have to work harder!"
	return m
}

func main() {
	sum := goFun.Add(5, 8)
	fmt.Println("The sum is", sum)

	greetings := goFun.Hello("Mwihoti");
	fmt.Println(greetings)


	age := 12
	if age >= 18 {
		fmt.Println("You a are now of legal age!")
	} else {
		fmt.Println("You are a minor!");
	}

	msg, legal  := is_legal_age(20)
	if legal {
		fmt.Println(msg)
	} else {
		fmt.Println(msg)
	}

	richMsg := is_rich(true)
	fmt.Println(richMsg)


}