package main

import "fmt"



func main() {
	fmt.Println(1 + 2)

	fmt.Println(5 * 1)

	n := 0
	n += 4
	fmt.Println(n)

	// fmt.Println(1 == 1)
	// fmt.Println(1 == 2)

	fmt.Println(true && false == true)
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)

	fmt.Println(true || false == true)
	fmt.Println(false || false == true)



}
