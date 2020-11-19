package computer

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func RunProgram(noun int, verb int) int {
	// read data
	data, err := ioutil.ReadFile("../data/input.txt")
	handleError(err)

	// turn input to array of ints instead of array of strings
	programStr := strings.Split(string(data), ",")
	var program = []int{}
	for _, str := range programStr {
		i, err := strconv.Atoi(str)
		handleError(err)
		program = append(program, i)
	}

	// insert noun and verb
	program[1] = noun
	program[2] = verb

	// run program
	programCounter := 0
	Loop:
		for programCounter < len(program) {
			switch program[programCounter] {
				case 1: // Add
					address1 := program[programCounter+1]
					address2 := program[programCounter+2]
					storeAddress := program[programCounter+3]
					result := program[address1] + program[address2]
					program[storeAddress] = result

					programCounter += 4

				case 2: 
					address1 := program[programCounter+1]
					address2 := program[programCounter+2]
					storeAddress := program[programCounter+3]
					result := program[address1] * program[address2]
					program[storeAddress] = result

					programCounter += 4

				case 99:
					break Loop

				default:
					log.Fatal("Unknown command '%s'\n", program[programCounter])
			}
	}
	return program[0]
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}