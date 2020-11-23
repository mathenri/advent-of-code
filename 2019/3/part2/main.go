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

  // follow and store path and distance of wire 1
  wire1visited := make(map[Pair]int64)
  startingPoint := Pair{0, 0}
  var startingDistance int64 = 0
  for _, step := range strings.Split(wire1, ",") {
    direction := string(step[0])
    stepsInt, err := strconv.Atoi(string(step[1:]))
    steps := int64(stepsInt)
    handleError(err)
    switch direction {
      case "U":
        var i int64
        for i = 1; i <= steps; i++ {
          wire1visited[Pair{startingPoint.X, startingPoint.Y+i}] = startingDistance+i
        }
        startingPoint = Pair{startingPoint.X, startingPoint.Y+steps}
        startingDistance += steps
      
      case "R":
        var i int64
        for i = 1; i <= steps; i++ {
          wire1visited[Pair{startingPoint.X+i, startingPoint.Y}] = startingDistance+i
        }
        startingPoint = Pair{startingPoint.X+steps, startingPoint.Y}
        startingDistance += steps
      
      case "D":
        var i int64
        for i = 1; i <= steps; i++ {
          wire1visited[Pair{startingPoint.X, startingPoint.Y-i}] = startingDistance+i
        }
        startingPoint = Pair{startingPoint.X, startingPoint.Y-steps}
        startingDistance += steps
      
      case "L":
        var i int64
        for i = 1; i <= steps; i++ {
          wire1visited[Pair{startingPoint.X-i, startingPoint.Y}] = startingDistance+i
        }
        startingPoint = Pair{startingPoint.X-steps, startingPoint.Y}
        startingDistance += steps
    }
  }

  // follow wire 2 and store intersections with wire 1
  intersections := make(map[Pair]int64)
  startingPoint = Pair{0, 0}
  var wire2distance int64 = 0
  for _, step := range strings.Split(wire2, ",") {
    direction := string(step[0])
    stepsInt, err := strconv.Atoi(string(step[1:]))
    steps := int64(stepsInt)
    handleError(err)
    switch direction {
      case "U":
        var i int64
        for i = 1; i <= steps; i++ {
          if wire1distance, ok := wire1visited[Pair{startingPoint.X, startingPoint.Y+i}]; ok {
            intersections[Pair{startingPoint.X, startingPoint.Y+i}] = wire1distance + wire2distance + i
          }
        }
        startingPoint = Pair{startingPoint.X, startingPoint.Y+steps}
        wire2distance += steps
      
      case "R":
        var i int64
        for i = 1; i <= steps; i++ {
          if wire1distance, ok := wire1visited[Pair{startingPoint.X+i, startingPoint.Y}]; ok {
            intersections[Pair{startingPoint.X+i, startingPoint.Y}] = wire1distance + wire2distance + i
          }
        }
        startingPoint = Pair{startingPoint.X+steps, startingPoint.Y}
        wire2distance += steps
      
      case "D":
        var i int64
        for i = 1; i <= steps; i++ {
          if wire1distance, ok := wire1visited[Pair{startingPoint.X, startingPoint.Y-i}]; ok {
            intersections[Pair{startingPoint.X, startingPoint.Y-i}] = wire1distance + wire2distance + i
          }
        }
        startingPoint = Pair{startingPoint.X, startingPoint.Y-steps}
        wire2distance += steps
      
      case "L":
        var i int64
        for i = 1; i <= steps; i++ {
          if wire1distance, ok := wire1visited[Pair{startingPoint.X-i, startingPoint.Y}]; ok {
            intersections[Pair{startingPoint.X-i, startingPoint.Y}] = wire1distance + wire2distance + i
          }
        }
        startingPoint = Pair{startingPoint.X-steps, startingPoint.Y}
        wire2distance += steps
    }
  }

  // calculate distance for each point
  var minDist int64 = math.MaxInt64
  minDistPoint := Pair{0, 0}
  for intersection, distance := range intersections {
    if distance < minDist {
      minDistPoint = intersection
      minDist = distance
    }
  }

  fmt.Printf("Minimum distance of a wire intersection to central port: %d (intersection %v)\n", minDist, minDistPoint)
}

func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}