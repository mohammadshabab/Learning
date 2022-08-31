package main

import "fmt"

func main() {
	ints := []int{1, 2, 3}
	strings := []string{"one", "two", "three"}
	Print(ints)
	Print(strings)
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
