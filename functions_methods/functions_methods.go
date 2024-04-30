package main

import "fmt"

func calcArea(r float64) float64 {
	return 3.14 * r * r
}
func calcPerimeter(r float64) float64 {
	return 2 * 3.14 * r
}
func calcDiameter(r float64) float64 {
	return 2 * r
}

func main() {
	fmt.Println("test line")
	var radius float64
	var query int
	fmt.Printf("Enter \n 1 - area \n 2 - perimeter \n 3 - diameter: \n") // Corrected typo
	fmt.Scanf("%d", &query)
	fmt.Print("Enter the value for the radius: ") // Corrected typo
	fmt.Scanf("%f", &radius)

	if query == 1 {
		fmt.Println("result:", calcArea(radius))
	} else if query == 2 {
		fmt.Println("result:", calcPerimeter(radius))
	} else if query == 3 {
		fmt.Println("result:", calcDiameter(radius))
	} else {
		fmt.Println("wrong entry..please check")
	}
	printResult(radius, getFunction(query))

	pointerFunc()
	pointerFunc1()
}

func printResult(radius float64, calcFunction func(r float64) float64) {
	result := calcFunction(radius)
	fmt.Println("result: ", result)
}
func getFunction(query int) func(r float64) float64 {
	query_to_func := map[int]func(r float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]

}

// A defer statement delays the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but
// the function call is not executed until the surrounding function returns.

//A pointer is a variable that holds memory address of another variable.
// They point where the memory is allocated and provide ways to find or even change the value located at the memory location.

//•& operator The address of a variable can be obtained by preceding the name of a variable with an ampersand sign (&), known as address of operator
//•* operator It is known as the dereference operator . When placed before an address, it returns the value at that address.

// declaring a pointer
// var <pointer_name> *<data_type>
// var ptr_i *int

func pointerFunc() {
	i := 10

	fmt.Printf("%T %v \n", &i, &i)       // value is : *int 0xc00000a138
	fmt.Printf("%T %v \n", *(&i), *(&i)) // int 10

}

func pointerFunc1() {
	var i *int

	fmt.Printf("%T %v \n", &i, &i)       // value is : **int 0xc000050028
	fmt.Printf("%T %v \n", *(&i), *(&i)) // *int <nil>

}
