package main

import (
    "fmt"
    "stackup.dev/intro-to-golang/goFun"
	"runtime"
	"log"

	"github.com/elastic/go-sysinfo"
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

func forLoop(c string) string {
	m := "Daniel"
	for i := 0; i < len(m); i++ {
		fmt.Println(string(m[i]))
	}
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


	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	
	default:
		fmt.Printf("%s.\n", os)
	}

	result := forLoop("Daniel")
	fmt.Println("forLoop result:", result)


	host, err := sysinfo.Host()
	if err != nil {
		log.Fatalln(err)
	} else {
		info := host.Info()
        fmt.Println("Hostname:", info.Hostname)
        fmt.Println("OS:", info.OS)
       
        fmt.Println("Kernel Version:", info.KernelVersion)
        fmt.Println("MAC Addresses:", info.MACs)
	}
}