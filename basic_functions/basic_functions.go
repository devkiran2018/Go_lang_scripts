package main

import (
	"fmt"
)

func main() {
	fmt.Println("basic Fucntions")
	// function format
	// func <function name> (<parameters>) <return tyep> {
	//body of the function or logic
	//}
	// Function arguments are the real values passed into the function.
	sumOfNumbers := addNumbers(5, 6)
	fmt.Println(sumOfNumbers)
	sum, diff := numbersOperation(10, 20)
	fmt.Println(sum, diff)
	sum1, diff1 := numbersOperation1(40, 50)
	fmt.Println(sum1, diff1)
	// sumNumebrs using variadic function
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(10, 20, 30))
	// variadic function
	printDetails("inetrn", "devops", "AI", "machineLearning", "Ops \n")

	// nested function
	x := func(length int, breadth int) int {
		return length * breadth
	}
	fmt.Printf("%T \n", x) //%T gives the type of X
	fmt.Println(x(10, 20))
	// decalre function with argumnets.
	y := func(length int, breadth int) int {
		return length * breadth
	}(20, 30)
	fmt.Printf("%T \n", y) //%T gives the type of X
	fmt.Println(y)
}

// Function parameters are the names listed in the function's definition.
func addNumbers(a int, b int) int { // return type is int and single return value.
	sum := a + b
	return sum
}
func numbersOperation(a int, b int) (int, int) { // return multiple values.
	sum := a + b // short hand reference of a+b result type is infered.
	diff := a - b
	return sum, diff
}

func numbersOperation1(a int, b int) (sum int, diff int) { // named return values.
	sum = a + b
	diff = a - b
	return // specified in defination.
}

// variadic functions  -> function that accepts variable number of arguments.
// t o declare a variadic function, the type of the final parameter is preceded by an ellipsis
// func <func_name >(param 1 type , param 2 type , para 3 ...type ) return_type
func sumNumbers(numbers ...int) int {
	sum := 0 // with in the function the numbers variable will have the slice conating all the arguments.
	// all the integer arguments passed will stored in a slice called numbers.
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func printDetails(student string, subjects ...string) {
	fmt.Println("hey", student, ",these are the trending subjects - ")
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
		// blank identifier '_' is the single udnerscore operator. it is used to ignore the values that are returned by the functions.
		//can be used as anonymous place holder.
	}
}

// An anonymous function is a function that is declared without any named identifier to refer to it.
// They can accept inputs and return outputs, just as standard functions do. used for short-term use.

// high order functions : function that receives a function as an argument or returns a function as output.
//Composition,  Reduces bugs, Code gets easier to read and understand.
