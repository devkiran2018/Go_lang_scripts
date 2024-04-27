package main

import (
	"fmt"
)

func main() {
	// array is the collection of similar data elements.
	// var <array name> [size of the array] <data type>
	var subjects [4]string
	subjects = [4]string{"one", "two", "three"}
	//fmt.Println(subjects)
	fmt.Println(len(subjects))
	fmt.Println(cap(subjects))

	for i := 0; i < len(subjects); i++ { // loop through an array
		fmt.Println(subjects[i])
	}
	fmt.Println(subjects[3]) // index 3 value is nil. so prints empty space.
	var subjectMarks [6]int
	fmt.Println("the marks in subjects are:", subjectMarks)
	for index, element := range subjects { //// loop through an array using range builtin
		fmt.Println(index, "=>", element)
	}

}
