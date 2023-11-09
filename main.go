package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	fmt.Println("Input c: ")
	fmt.Scan(&c)

	fmt.Println("V = ", a*b*c)

	fmt.Println("S = ", 2*((a*b)+(b*c)+(a*c)))
}
