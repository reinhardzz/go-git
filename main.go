package main

import "fmt"


func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello Github")
	fmt.Println(sum(1, 2))
	fmt.Println(sub(1, 2))
	fmt.Println(mul(1, 2))
}
