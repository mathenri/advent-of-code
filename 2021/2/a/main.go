package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	// read input file
	inputFile, err := os.Open("../input.txt")
  if err != nil {
      panic(err)
  }
  defer inputFile.Close()

  // parse input file
  horizontalPos := 0
  depth := 0
  scanner := bufio.NewScanner(inputFile)
  for scanner.Scan() {

  	// parse input command
  	row := scanner.Text()
  	tokens := strings.Split(row, " ")
  	if len(tokens) != 2 {
  		fmt.Println("invalid input: ", row)
  		continue
  	}
  	cmd := tokens[0]
  	val, err := strconv.Atoi(tokens[1])
  	if err != nil {
  		panic(err)
  	}

  	// execute command
  	switch cmd {
  	case "forward":
  		horizontalPos += val
  	case "down":
  		depth += val
  	case "up":
  		depth -= val
  	default:
  		fmt.Println("Unknown command:", cmd)
  		return
  	}
  }

  if err := scanner.Err(); err != nil {
      panic(err)
  }
  
  fmt.Println("Result:", horizontalPos * depth)
}