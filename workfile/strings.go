package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Hello\n")

	// This is string
	var ourString string
	ourString = "Hello Go world\n"

	fmt.Print(ourString)

	// This is runego run strings.go
	var ourRune rune
	ourRune = 'G'
	fmt.Printf("type:%T, \n value:%v\n", ourRune, ourRune)

	//Example of needing special characters
	fmt.Print("Don't touch that,that's my food\n")
	fmt.Print("She said: \"that's my food\"")

	//Example of string literal
	fmt.Println("\n", `This is my practise on the GO
	and I am loving its`)

	//Example of length of a character
	name := "Lateef Emmanuel Babatunde"

	fmt.Println(len(name))

	//string formatting
	nameOne := "Emmanuel"
	num := 1000

	fmt.Printf("my name is %v and my money is %v", nameOne, num)

	//iteration

	number := []int{12, 45, 60, 43, 38}
	for _, i := range number {
		fmt.Println(i)
	}

	//use of string method

	fN := "Emmanuel"
	fNSplit := strings.Split(fN, "l")
	fmt.Println(fNSplit)
}
