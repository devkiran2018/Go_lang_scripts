package main

import (
	"fmt"
)

const PI float64 = 3.142 // global constant
func main() {

	// test inner and outer block of code
	your_name := "testName"
	{
		dob := "secret"
		fmt.Println("should only be accessed form with in the inner block:", dob)
		fmt.Println("what is your name:", your_name) // accessing from outer block
	}
	fmt.Println("what is your name", your_name)
	//fmt.Println("whai is your dob", dob) // throw error when acccessed inner block variable

	fmt.Println("test line")
	for i := 0; i < 5; i++ { //initialization ; condition ; post
		if i == 3 { //!=, <=, >=
			continue // skips the current iteration
		} else if i == 4 {
			break // ends the loop immediately when condition met.
		} else {
			fmt.Printf("the value of i is %d \n", i)
		}
	}
	i := 1 // another form of for loop
	for i <= 3 {
		fmt.Println("value of i is:", i)
		i += 1
	}
	var name string = "legal"
	switch name {
	case "not lehgal":
		fmt.Println("name is not legal")

	case "legal":
		fmt.Println("it is legal name")
		fallthrough // forces the execution flow to fall through other case blocks.
	default:
		fmt.Println(" given name is neither of the case")
	}
}
