package main

import "fmt"

type User struct {
	Name string
}

// BAD: returns pointer to heap-allocated struct
func newUser(name string) *User {
	return &User{Name: name} // escapes to heap
}

// BETTER: use value types if pointer is unnecessary
func printUser(u User) {
	fmt.Println(u.Name)
}

func main() {
	_ = newUser("roma")
	printUser(User{"ROMA"})
}
