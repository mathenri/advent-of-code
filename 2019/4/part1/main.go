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
	containsDoubleNumber := false
	for i, c := range password {
		if i == 0 {
			continue // skip first number
		}

		previousNumber := string(password[i-1])

		if string(c) == previousNumber {
			containsDoubleNumber = true
		}

		if previousNumber > string(c) {
			return false // password not always increasing
		}
	}
	return containsDoubleNumber
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}