//go:build wasm
// +build wasm

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
}

//go:export sum
func sum(a, b int) int {
	return a + b
}
