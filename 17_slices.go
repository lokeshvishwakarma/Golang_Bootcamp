package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 4}
	b := a // 'a' and 'b' are pointing to the same block of slice
	b[1] = 5
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("Length: %v, %T\n", len(a), len(a))
	fmt.Printf("Capacity: %v, %T\n", cap(a), cap(a)) // capacity of a slice
	fmt.Printf("-------------------------\n")

	// l := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // array will also work same but assignment will create new copies
	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := l[:]   // slice of all the elements
	n := l[3:]  // slice from 4th element to end
	o := l[:6]  // slice first 6 elements
	p := l[3:6] // slice the 4th, 5th and 6th elements
	l[5] = 42   // slices point to the same block of memory
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", o, o)
	fmt.Printf("%v, %T\n", p, p)
	fmt.Printf("-------------------------\n")
}
