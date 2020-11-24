package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// get input
	fromStr := os.Args[1]
	toStr := os.Args[2]
	from, err := strconv.Atoi(fromStr)
	handleError(err)
	to, err := strconv.Atoi(toStr)
	handleError(err)

	nrValidPasswords := 0
	for i := from; i <= to; i++ {
		password := strconv.Itoa(i)
		handleError(err)
		if isValidPassword(password) {
			nrValidPasswords++;
		}
	}

	fmt.Println("Number of valid passwords: ", nrValidPasswords)
}

func isValidPassword(password string) bool {
	charCounts := make(map[rune]int) 
	for i, c := range password {
		
		// count chars
		if _, ok := charCounts[c]; ok {
			charCounts[c]++
		} else {
			charCounts[c] = 1
		}

		if i == 0 {
			continue
		}

		// check if numbers are not increasing
		previousNumber := string(password[i-1])
		if previousNumber > string(c) {
			return false // password not always increasing
		}
	}

	containsDouble := false
	for _, count := range charCounts {
		if count == 2 {
			containsDouble = true
		}
	}
	return containsDouble
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}