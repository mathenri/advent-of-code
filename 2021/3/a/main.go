package main

import (
	"fmt"
	"bufio"
	"os"
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
  oneCounts := [12]int{}
  inputNrRows := 0
  scanner := bufio.NewScanner(inputFile)
  for scanner.Scan() {
  	inputNrRows++
  	for i := 0; i < 12; i++ {
  		num := scanner.Text()
  		bit := string(num[i])
  		switch bit {
  		case "1":
  			oneCounts[i]++
  		case "0":
  			// do nothing
  		default:
  			fmt.Println("Invalid bit value:", bit)
  		}
  	}
  }

  epsilonStr := ""
  gammaStr := ""
  for i := 0; i < 12; i++ {
  	fmt.Println("i", i)
  	fmt.Println("oneCounts[i]", oneCounts[i])
  	if oneCounts[i] > (inputNrRows/2) {
  		gammaStr += "1"
  		epsilonStr += "0"
  	} else {
  		gammaStr += "0"
  		epsilonStr += "1"
  	}
  }

  epsilon, err := strconv.ParseInt(epsilonStr, 2, 64)
  if err != nil {
  	panic(err)
  }

  gamma, err := strconv.ParseInt(gammaStr, 2, 64)
  if err != nil {
  	panic(err)
  }

  fmt.Println("Result:", epsilon * gamma)
  
}