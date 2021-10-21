package main

import "fmt"

//go:generate go run internal/generate/generatefoobar.go

func main() {
	fmt.Println("hello")
}
