package main

// import (
// 	"fmt"
// )

// const (
// 	isAdmin = 1 << iota
// 	isHeadQuarter
// 	canSeeFinancials

// 	canSeeAfrica
// 	canSeeAsia
// 	canSeeEurope
// 	canSeeNorthAmerica
// 	canSeeSouthAmerica
// )

// func main() {
// 	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
// 	fmt.Printf("%b\n", isAdmin)
// 	fmt.Printf("%b\n", roles)
// 	fmt.Printf("Is Admin? %v\n", isAdmin&roles)
// 	fmt.Printf("Is Headquarter? %v\n", isHeadQuarter&roles)
// 	fmt.Printf("-------------------------\n")
// }

// Constants are immutable but shadowed
// Can interoperate only with same type
// PascalCase for exported constants
// camelCase for internal constants
