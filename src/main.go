package main

import "fmt"

func main() {
	var firstName *string

	*firstName = "Arthur"

	fmt.Println(firstName)
}
