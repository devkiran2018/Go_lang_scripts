package main

import (
	"fmt"
)

func main() {
	fmt.Println("test line")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break

		}
		fmt.Printf("the value of i is %d \n", i)
	}

}
