package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	foo := []int{1, 1, 2, 2, 3}

	fmt.Println(topKFrequent(foo, 2))
}
