package main

import (
	"fmt" // format package for all print functionality.
)

// where to start the program? where is the entry point? -> starts from mian() function
// this is online comment

/* multiline comment
this is comment on second line.
*/

func main() {
	//var <variablename> <datatype> = <value>
	var variableName1 = "Go data types" //explicitly declared avriable.
	variableName2 := "test data"        // implcit that means go will infer the type of the varibale when := is used.
	// %s plain string, %v default format, %T type of the value, %d inetegrs, %f flaoting numbers, %.2f floating up to 2 decimal places.

	fmt.Print("this will print:", variableName1, "\n")
	fmt.Printf("this will print as %T %s \n", variableName1, variableName1)
	fmt.Println("Println does not need any new line character at the end", variableName2)

}
