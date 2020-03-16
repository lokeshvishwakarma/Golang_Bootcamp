package main

// import (
// 	"fmt"
// )

// func main() {
// 	const a int = 10
// 	// a = 20		// cannot change the values of constants
// 	fmt.Printf("%v, %T\n", a, a)
// 	fmt.Printf("-------------------------\n")
// 	// const myconst = math.Sin(1.57)		// cannot assign values to constants at runtime
// 	// fmt.Printf("%v, %T\n", myconst, myconst)
// 	const (
// 		b = iota // used to create enumerated constants
// 		c = iota // used to create enumerated constants
// 		d = iota // used to create enumerated constants
// 	)
// 	fmt.Printf("%v, %T\n", b, b)
// 	fmt.Printf("%v, %T\n", c, c)
// 	fmt.Printf("%v, %T\n", d, d)
// 	fmt.Printf("-------------------------\n")

// 	const (
// 		l = iota // used to create enumerated constants
// 		m        //not necessary to assign iota in single constant block
// 		n
// 	)
// 	fmt.Printf("%v, %T\n", l, l)
// 	fmt.Printf("%v, %T\n", m, m)
// 	fmt.Printf("%v, %T\n", n, n)
// 	fmt.Printf("-------------------------\n")

// 	const x = iota // for enumeration the constants have to be in the same block
// 	const y = iota
// 	const z = iota
// 	fmt.Printf("%v, %T\n", x, x)
// 	fmt.Printf("%v, %T\n", y, y)
// 	fmt.Printf("%v, %T\n", z, z)
// }
