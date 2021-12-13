package main

import (
	"fmt"
	"bufio"
	"os"
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
  parenthesisCount := 0
  squareBracketCount := 0
  curlyBracketCount := 0
  ltSignCount := 0
  for scanner.Scan() {
    line := scanner.Text()
    incorrectClosingChar := firstIncorrectClosingChar(line)
    switch incorrectClosingChar {
    case ")":
  		parenthesisCount++
  	case "]":
  		squareBracketCount++
  	case "}":
  		curlyBracketCount++
  	case ">":
  		ltSignCount++
  	}
  }

  if err := scanner.Err(); err != nil {
      panic(err)
  }

  fmt.Println("Result", parenthesisCount*3 + squareBracketCount*57 + curlyBracketCount*1197 + ltSignCount*25137)
}

func firstIncorrectClosingChar(inputStr string) string {
	openBracketStack := Stack{}
  for _, byteC := range inputStr {
  	c := string(byteC)
  	switch c {
  	case "(", "[", "{", "<":
  		openBracketStack.Push(c)
  	case ")", "]", "}", ">":
  		lastOpenBracket := openBracketStack.Pop()
  		if !bracketMatch(lastOpenBracket, c) {
  			return c
  		}  		
  	}
  }
  return ""
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