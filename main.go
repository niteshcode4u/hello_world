package main

import "fmt"

// GetMessage returns the "Hello, World!" message
func GetMessage() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(GetMessage())
}
