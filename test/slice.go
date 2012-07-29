package main

import "fmt"

// Both length and capacity (i.e. the largest amount the underlying array can grow.)

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])
}
