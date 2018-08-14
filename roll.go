package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// MinStringLength is the minimum length for a dice string expression, e.g. "1d1"
const MinStringLength = 3

func main() {
	// Make sure we're as close to random as possible by setting the seed
	var seed = time.Now().UnixNano()
	rand.Seed(seed)

	var args []string
	// The os.Args always contains the path to the executable so program args are at index 1+
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	var dice string
	// If we have an arg and it matches the min length, use it
	if args != nil && len(args[0]) >= MinStringLength {
		dice = args[0]
	} else {
		// Otherwise prompt the user for a string
		dice = getInput()
	}

	// if len(dice) >= MinStringLength {
	if isValidDiceString(dice) {
		roller(dice)
	} else {
		fmt.Println("Couldn't understand", dice)
	}
}

func roller(dice string) {
	var parsedDice = strings.Split(dice, "d")
	fmt.Println("Rolling", dice)
	var total = 0

	if count, err := strconv.Atoi(parsedDice[0]); err != nil {
		fmt.Printf("Error 1: %+v", err)
	} else {
		if numSides, err := strconv.Atoi(parsedDice[1]); err != nil {
			fmt.Printf("Error 2: %+v", err)
		} else {
			i := 1
			for i <= count {
				var rolledValue = rand.Intn(numSides) + 1
				total += rolledValue
				fmt.Println(i, "of", count, "rolled", rolledValue)
				i++
			}
		}
	}
	fmt.Println("Total rolled was", total)
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Roll: ")
	dice, _ := reader.ReadString('\n')
	dice = strings.TrimSuffix(dice, "\n")
	return dice
}

func isValidDiceString(dice string) bool {
	if len(dice) < MinStringLength {
		return false
	}
	// This regex matches the format "1d1" exactly
	return regexp.MustCompile("^([0-9])+d([0-9])+$").MatchString(dice)
}
