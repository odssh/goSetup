package main

import "fmt"

func main() {
	var name string = "ods"
	var age uint8
	age = 21
	var ability int8
	ability = 127
	company := true
	var potential uint8
	var iq bool
	fmt.Println("Hello", name, age, ability, company, potential, iq)
	fmt.Printf("Hey %T from %v", name, company)
	var x string = fmt.Sprintf("\nHey %T from %v", name, iq)
	fmt.Println(x)
}
