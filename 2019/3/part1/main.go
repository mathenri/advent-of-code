package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"math"
)

type Pair struct {
	X, Y int64
}

func main() {
	// read input
	data, err := ioutil.ReadFile("input.txt")
	handleError(err)

	// parse input
	wires := strings.Split(string(data), "\n")
	wire1 := wires[0]
	wire2 := wires[1]

	// follow and store path of wire 1
	wire1visited := make(map[Pair]bool)
	startingPoint := Pair{0, 0}
	for _, step := range strings.Split(wire1, ",") {
		direction := string(step[0])
		stepsInt, err := strconv.Atoi(string(step[1:]))
		steps := int64(stepsInt)
		handleError(err)
		switch direction {
			case "U":
				for y := startingPoint.Y+1; y <= startingPoint.Y+steps; y++ {
					wire1visited[Pair{startingPoint.X, y}] = true
				}
				startingPoint = Pair{startingPoint.X, startingPoint.Y+steps}
			
			case "R":
				for x := startingPoint.X+1; x <= startingPoint.X+steps; x++ {
					wire1visited[Pair{x, startingPoint.Y}] = true
				}
				startingPoint = Pair{startingPoint.X+steps, startingPoint.Y}
			
			case "D":
				for y := startingPoint.Y-1; y >= startingPoint.Y-steps; y-- {
					wire1visited[Pair{startingPoint.X, y}] = true
				}
				startingPoint = Pair{startingPoint.X, startingPoint.Y-steps}
			
			case "L":
				for x := startingPoint.X-1; x >= startingPoint.X-steps; x-- {
					wire1visited[Pair{x, startingPoint.Y}] = true
				}
				startingPoint = Pair{startingPoint.X-steps, startingPoint.Y}
		}
	}

	// follow wire 2 and store intersections with wire 1
	intersections := []Pair{}
	startingPoint = Pair{0, 0}
	for _, step := range strings.Split(wire2, ",") {
		direction := string(step[0])
		stepsInt, err := strconv.Atoi(string(step[1:]))
		steps := int64(stepsInt)
		handleError(err)
		switch direction {
			case "U":
				for y := startingPoint.Y+1; y <= startingPoint.Y+steps; y++ {
					if wire1visited[Pair{startingPoint.X, y}] {
						intersections = append(intersections, Pair{startingPoint.X, y})
					}
				}
				startingPoint = Pair{startingPoint.X, startingPoint.Y+steps}
			
			case "R":
				for x := startingPoint.X+1; x <= startingPoint.X+steps; x++ {
					if wire1visited[Pair{x, startingPoint.Y}] {
						intersections = append(intersections, Pair{x, startingPoint.Y})
					}
				}
				startingPoint = Pair{startingPoint.X+steps, startingPoint.Y}
			
			case "D":
				for y := startingPoint.Y-1; y >= startingPoint.Y-steps; y-- {
					if wire1visited[Pair{startingPoint.X, y}] {
						intersections = append(intersections, Pair{startingPoint.X, y})
					}
				}
				startingPoint = Pair{startingPoint.X, startingPoint.Y-steps}
			
			case "L":
				for x := startingPoint.X-1; x >= startingPoint.X-steps; x-- {
					if wire1visited[Pair{x, startingPoint.Y}] {
						intersections = append(intersections, Pair{x, startingPoint.Y})
					}
				}
				startingPoint = Pair{startingPoint.X-steps, startingPoint.Y}
		}
	}

	// calculate manhattan distance for each point
	var minDist int64 = math.MaxInt64
	minDistPoint := Pair{0, 0}
	for _, intersection := range intersections {
		dist := manhattanDistanceToOrigo(intersection.X, intersection.Y)
		if dist < minDist {
			minDistPoint = intersection
			minDist = dist
		}
	}

	fmt.Printf("Minimum distance of a wire intersection to central port: %d (intersection %v)", minDist, minDistPoint)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func manhattanDistanceToOrigo(x int64, y int64) int64 {
	return abs(x) + abs(y)
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}