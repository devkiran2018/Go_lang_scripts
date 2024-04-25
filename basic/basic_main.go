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
	var variableName1 = "Go data types"
	variableName2 := "test data"

	fmt.Printf("this will print %T \n", variableName1)
	fmt.Println("Println does not need any '\n' at the end %s ", variableName2)
}
