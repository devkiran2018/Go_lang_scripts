package main

import (
	"fmt" // format package for all print functionality.
	"reflect"
	"strconv"
)

// where to start the program? where is the entry point? -> starts from mian() function
// this is online comment

/* multiline comment
this is comment on second line.
*/

func main() {
	fmt.Println("the sum of 2 numbers is", add(10, 20)) // calling a function
	//var <variablename> <datatype> = <value>
	var variableName1 = "Go data types" //explicitly declared avriable.
	variableName2 := "test data"        // implcit that means go will infer the type of the varibale when := is used.
	// %s plain string, %v default format, %T type of the value, %d inetegrs, %f flaoting numbers, %.2f floating up to 2 decimal places.
	var userName string   // declare a variable
	userName = "testName" // assign value to varibale.
	userName = "testName1"
	//short hand var decleartion
	var a, b string = "var1", "var2"
	// declare diffrent types in 1 go.
	var (
		a1 string = "var3"
		b1 int    = 5
		//a1 string
		//b1 int
	) // zero values if values are not assigned-> bool false, int 0, float64 0.0, string ""
	fmt.Print("this will print:", userName, a, b, "\n")
	fmt.Print("this will print:", a1, "-", b1, "\n")
	fmt.Printf("this will print as %T %s \n", variableName1, variableName1)
	fmt.Println("Println does not need any new line character at the end", variableName2)
	var name string
	//fmt.Scanf("%s", &name)
	fmt.Print("this will print scanf ", name, "\n")
	fmt.Print("input the values for a_input,  binput:", "\n")
	// var (
	// 	a_inputString string
	// 	b_inputNum    int
	// )
	// fmt.Scanf(" %s %d", &a_inputString, &b_inputNum)
	// fmt.Print("this will print scanned values as ", "\n")
	//type of variable %T
	var name_type string = "test1"
	fmt.Printf("name_type= %v is of type %T", name_type, name_type)
	fmt.Printf("type: %v \n", reflect.TypeOf(10))

	var i int = 10
	var f float64 = float64(i)

	fmt.Printf("%.2f \n", f) // outputs as 10.00
	// string conversion package for string to int or int to string
	// Atoi(), Itoa()
	var s string = strconv.Itoa(i)
	fmt.Printf("%q \n", s)
	s = "10a"
	i, err := strconv.Atoi(s)
	fmt.Printf("default value for int is %v,%T\n", i, i)
	fmt.Printf("this is the %v, %T \n", err, err)

}

func add(a int, b int) int {
	return a + b

}
