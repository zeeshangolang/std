package main

import (
	"fmt"
)

//four families of printing
//Print, Println, Printf -> os.Stdout
//Sprint, Sprintln and Sprint returns a string

func main() {

	const name, age = "zeeshan", 21
	//PRINT
	fmt.Print(name, " is ", age, " years old.\n")
	//PRINTF
	fmt.Printf("%s is %d years old\n", name, age)
	//PRINTLN
	fmt.Println(name, "is", age, "years old")
	fmt.Println(name, "is", age, "years old")

	fmt.Println("hello")

}
