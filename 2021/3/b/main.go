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

  // input file to memory buffer
  allRows := []string{}
  scanner := bufio.NewScanner(inputFile)
  for scanner.Scan() {
  	allRows = append(allRows, scanner.Text())
  }
  
  oxygenGeneratorRatingStr := findOxygenGeneratorRating(allRows)
  co2ScrubberRatingStr := findCO2ScrubberRating(allRows)

  oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatingStr, 2, 64)
  if err != nil {
  	panic(err)
  }
  co2ScrubberRating, err := strconv.ParseInt(co2ScrubberRatingStr, 2, 64)
  if err != nil {
  	panic(err)
  }

  fmt.Println("Result:", oxygenGeneratorRating * co2ScrubberRating)
  
}

func findOxygenGeneratorRating(inputRows []string) string {
	return findMetric(
		inputRows,
		func(nrOnes, totalNrRows int) bool {
			return float64(nrOnes) >= (float64(totalNrRows)/2)
		},
	)
}

func findCO2ScrubberRating(inputRows []string) string {
	return findMetric(
		inputRows,
		func(nrOnes, totalNrRows int) bool {
			return float64(nrOnes) < (float64(totalNrRows)/2)
		},
	)
}

func findMetric(inputRows []string, selectOnesFunc func(nrOnes, totalNrRows int) bool) string {
	previousNumbersSelection := inputRows
  for bitIndex := 0; bitIndex < 12; bitIndex++ {
  	inputNrRows := 0
  	nrOnes := 0
  	targetNumber := "0"
  	newNumbersSelection := []string{}
  	
  	for _, row := range previousNumbersSelection {
  		inputNrRows++
  		if string(row[bitIndex]) == "1" {
  			nrOnes++
  		}
  	}
  	
  	if selectOnesFunc(nrOnes, inputNrRows) {
  		targetNumber = "1"
  	}
  	
  	for _, row := range previousNumbersSelection {
  		if string(row[bitIndex]) == targetNumber {
  			newNumbersSelection = append(newNumbersSelection, row)
  		}
  	}
  	
  	if len(newNumbersSelection) == 1 {
  		return newNumbersSelection[0]
  	} else {
  		previousNumbersSelection = newNumbersSelection
  	}
  }
  panic("never found the final number")
} 