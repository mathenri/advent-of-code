package main

import (
  computer "../common"
  "fmt"
)

func main() {
	Loop:
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				result := computer.RunProgram(noun, verb)
				fmt.Printf("Running with noun '%d' and verb '%d' gave result '%d'\n", noun, verb, result)
				if result == 19690720 {
					fmt.Printf("Found answer!\n")
					break Loop
				}
			}
		}
}