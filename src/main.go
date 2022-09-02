package main

import "fmt"

func main() {
	type user struct {
		Id        int
		FirstName string
		LastName  string
	}

	var u user
	u.Id = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"

	fmt.Println(u.FirstName)

	u2 := user{Id: 1, FirstName: "Arthur", LastName: "Dent"}

	fmt.Println(u2)
}
