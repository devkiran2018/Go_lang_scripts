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
	for _, value := range subjects { // _ => empty index, only values are printed.
		fmt.Println(value)
	}
	// multidimensional arrays
	arr := [2][3]int{{2, 4, 6}, {4, 6, 8}}
	fmt.Println(arr[0][1])
	fmt.Println(arr[0][2])
	dataSlice() // calling function for slice
	fmt.Println("calling slice")
	makeSlice() // make slice
	fmt.Println("this is from append slice func:")
	appendSlice()
	fmt.Println("this is from mapDatatype() func:")
	mapsDatatype()
}
func dataSlice() {
	// continuous segment of an underlying array. // varibale typed, more flexible
	// <slice_name> := []<data_type> {<values>}

	subjectsMarksArr := [6]int{60, 40, 50, 80, 90, 100} //array
	fmt.Println("the cap of array", cap(subjectsMarksArr))
	slice_1 := subjectsMarksArr[1:4]
	fmt.Println("the cap of slice", cap(slice_1))
	sub_slice := slice_1[0:2]
	fmt.Println(sub_slice)
}
func makeSlice() {
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
func appendSlice() {
	arr := [4]int{10, 20, 30, 40}
	slice1 := arr[1:3]
	fmt.Println(slice1)
	slice1 = append(slice1, 60, 70) // slice= append (slice, anotherSlice...)
	fmt.Println(slice1)             // append to existing slice
	slice2 := arr[0:1]
	new_slice := append(slice1, slice2...)
	fmt.Println(new_slice)

	i := 2
	fmt.Println(arr)
	slice_1 := arr[:i]
	slice_2 := arr[i+1:]
	new_slice1 := append(slice_1, slice_2...) // element index at 2 ie is 30 is removed form the new slice
	fmt.Println(new_slice1)
	//copying from a slice
	src_slice := []int{10, 20, 30, 40, 50, 60}
	dest_slice := make([]int, 3)
	num := copy(dest_slice, src_slice)
	fmt.Println(dest_slice)
	fmt.Println("the numebr of elements copied:", num)

}
func mapsDatatype() {
	// unordered collection of kye-value pairs, implemmeted by hash table.
	// var map_name map[key_data_type]<value_data_type>
	// var my_map map[string]string
	var codes map[string]string //this syntax creates a nil map.
	// code["one"] = "1"  this will throw and error due, nil map, which does not have key.
	// maps needs to be initialised in different way.
	// maps needs to declared and initialized
	fmt.Println(codes) // := shorthand operator
	// map_name := map[key_data_type]value_data_type{key_value_pairs}
	codes_map := map[string]string{"one": "1", "two": "2"}
	// add key-value pair
	codes_map["three"] = "3"
	codes_map["four"] = "4"
	// update
	codes_map["one"] = "one1"
	fmt.Println(codes_map)
	// delete key-value-pair
	delete(codes_map, "four")
	fmt.Println(codes_map["one"]) //accessing map with keys.
	value, found := codes_map["two"]
	fmt.Println(found, value)
	// ietrate over the map
	for key, value := range codes_map {
		fmt.Println(key, "==>", value)
	}
	// declare and initilize a map with make() function
	new_map := make(map[string]string)
	fmt.Println(new_map) // empty map[]
	// truncate a map using delete or make function
	//for key, value := range codes_map {
	//	delete(code_map,key)
	//}
	codes_map = make(map[string]string)
	fmt.Println(codes_map)

}
