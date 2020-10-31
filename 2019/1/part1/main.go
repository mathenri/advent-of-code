package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"log"
)

func main() {
	
	// read input file
	data, err := ioutil.ReadFile("input.txt")
	handleError(err)
	
	// calculate fuel needed
	var totalFuel float64 = 0
	for _, line := range strings.Split(string(data), "\n") {
		mass, err := strconv.Atoi(line)
		handleError(err)
		fuel := math.Floor(float64(mass / 3.0)) - 2
		totalFuel += fuel
	}
	fmt.Printf("Total fuel %f", totalFuel)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}