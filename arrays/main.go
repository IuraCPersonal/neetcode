package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	foo := []int{-1, 1, 0, -3, 3}

	fmt.Println(productExceptSelf(foo))
}
