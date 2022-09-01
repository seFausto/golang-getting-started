package main

import "fmt"

const (
	first = iota + 6
	second
)

const (
	third = iota
)

func main() {
	fmt.Println(first, second, third)
}
