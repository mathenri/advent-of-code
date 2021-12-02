package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	// read input file
	inputFile, err := os.Open("input.txt")
  if err != nil {
      panic(err)
  }
  defer inputFile.Close()

  // parse input file
  scanner := bufio.NewScanner(inputFile)

  // get first depth
  scanner.Scan()
  firstDepthStr := scanner.Text()
  prevDepth, err := strconv.Atoi(firstDepthStr)
  if err != nil {
      panic(err)
  }

  // count increasing depths
  depthIncreses := 0
  for scanner.Scan() {
    depthStr := scanner.Text()
    depth, err := strconv.Atoi(depthStr)
	  if err != nil {
	      panic(err)
	  }
	  if depth > prevDepth {
	  	depthIncreses++
	  }
	  prevDepth = depth
  }

  if err := scanner.Err(); err != nil {
      panic(err)
  }

  fmt.Printf("Depth increased %d times\n", depthIncreses)
}