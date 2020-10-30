package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "math"

func main() {
	
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		var totalFuel float64 = 0
		for _, line := range strings.Split(string(data), "\n") {
			mass, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			} else {
				fuel := math.Floor(float64(mass / 3.0)) - 2
				totalFuel += fuel
			}
		}
		fmt.Printf("Total fuel %f", totalFuel)
	}
}