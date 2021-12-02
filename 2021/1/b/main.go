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


  // count increasing depths
  var depths []int
  depthIncreses := 0
  i := 0
  for scanner.Scan() {
      
		// get single depth and add to buffer
    depthStr := scanner.Text()
    depth, err := strconv.Atoi(depthStr)
	  if err != nil {
	      panic(err)
	  }
	  depths = append(depths, depth)

	  // we need at least 4 depths
	  if len(depths) < 4 {
	  	i++
	  	continue
	  }

	  prevDepth := depths[i-3] + depths[i-2] + depths[i-1]
	  currentDepth := depths[i-2] + depths[i-1] + depths[i]

	  if currentDepth > prevDepth {
	  	depthIncreses++
	  }

	  i++
  }

  if err := scanner.Err(); err != nil {
      panic(err)
  }

  fmt.Printf("Depth increased %d times\n", depthIncreses)
}