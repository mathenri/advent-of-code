package main

import (
  computer "../common"
  "fmt"
)

func main() {
	result := computer.RunProgram(12, 2)
	fmt.Printf("Result at position 0: %d\n", result)
}