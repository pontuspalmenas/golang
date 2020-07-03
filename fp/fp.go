package main

import "fmt"

func do(f func(a int, b int) int, a int, b int) int {
	return f(a, b)
}

func main() {
	add := func(a int, b int) int {
		return a + b
	}

	sub := func(a int, b int) int {
		return a - b
	}

	fmt.Println(do(add, 1, 2))
	fmt.Println(do(sub, 1, 2))
}
