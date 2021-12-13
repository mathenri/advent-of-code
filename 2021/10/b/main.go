package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"errors"
)

func main() {
	
	// read input file
	inputFile, err := os.Open("../input.txt")
  if err != nil {
      panic(err)
  }
  defer inputFile.Close()

  // parse input file
  scanner := bufio.NewScanner(inputFile)
  closingBracketScores := []int{}
  for scanner.Scan() {
    line := scanner.Text()
    missingClosingBracketsSequence, err := getMissingClosingBrackets(line)
    if err != nil {
    	continue
    }

    score := getClosingBracketScore(missingClosingBracketsSequence)
    closingBracketScores = append(closingBracketScores, score)
  }

  if err := scanner.Err(); err != nil {
      panic(err)
  }

  sort.Ints(closingBracketScores)
  middleIndex := (len(closingBracketScores) / 2)
  middleScore := closingBracketScores[middleIndex]

  fmt.Println("Result", middleScore)
}

func getClosingBracketScore(closingBracketSequence string) int {
	score := 0
	for _, byteC := range closingBracketSequence {
		score *= 5
		c := string(byteC)
		switch c {
		case ")":
			score += 1
		case "]":
			score += 2
		case "}":
			score += 3
		case ">":
			score += 4
		}
	}
	return score
}

func getMissingClosingBrackets(inputStr string) (string, error) {
	openBracketStack := Stack{}
  for _, byteC := range inputStr {
  	c := string(byteC)
  	switch c {
  	case "(", "[", "{", "<":
  		openBracketStack.Push(c)
  	case ")", "]", "}", ">":
  		lastOpenBracket := openBracketStack.Pop()
  		if !bracketMatch(lastOpenBracket, c) {
  			return "", incorrectClosingBracketError
  		}  		
  	}
  }

  missingClosingBracketsSequence := ""
  for openBracketStack.Len() != 0 {
  	openBracket := openBracketStack.Pop()
  	switch openBracket {
  	case "(":
  		missingClosingBracketsSequence += ")"
  	case "[":
  		missingClosingBracketsSequence += "]"
  	case "{":
  		missingClosingBracketsSequence += "}"
  	case "<":
  		missingClosingBracketsSequence += ">"
  	}

  }
  return missingClosingBracketsSequence, nil
}

func bracketMatch(open, close string) bool {
	if open == "(" && close == ")" {
		return true
	}
	if open == "[" && close == "]" {
		return true
	}
	if open == "{" && close == "}" {
		return true
	}
	if open == "<" && close == ">" {
		return true
	}
	return false
}

type Stack struct {
	arr []string
}

func (s *Stack) Push(e string) {
	s.arr = append(s.arr, e)
}

func (s *Stack) Pop() string {
	lastElem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return lastElem
}

func (s *Stack) Len() int {
	return len(s.arr)
}
 
var incorrectClosingBracketError error = errors.New("Incorrect closing bracket")