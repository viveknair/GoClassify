package main

import(
	"fmt"
)

func testing() ([]int) {
	slicingTest := make([]int,10)
	slicingTest[0] = 5
	return slicingTest
}

func main() {
	slicingTest := testing()
	fmt.Printf("The test is %v \n", slicingTest[0] )
}
