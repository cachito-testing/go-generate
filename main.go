package main

import (
	"fmt"
	"rsc.io/quote"
)

//go:generate go run internal/generate/generatefoobar.go

func main() {
	fmt.Println(quote.Hello())
}
