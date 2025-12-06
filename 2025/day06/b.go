package main

import (
	"bufio"
	"fmt"
	"github.com/teoc98/advent-of-code/2025/utils"
	"log"
	"os"
	// "regexp"
	"strconv"
)

func operate(operator string, operands []int) int {
	var value int
	var fn func(int, int) int

	if operator == "+" {
		value = 0
		fn = func(x, y int) int {
			return x + y
		}
	} else if operator == "*" {
		value = 1
		fn = func(x, y int) int {
			return x * y
		}
	}

	for i := 0; i < len(operands); i += 1 {
		value = fn(value, operands[i])
	}

	return value
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments")
	}
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var worksheet []string
	grandTotal := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		worksheet = utils.SliceAppend(worksheet, line)
	}

	m := len(worksheet)
	n := len(worksheet[0])
	var operands []int
	var operator string

	for i := 0; i < n; i += 1 {
		allSpaces := true

		operand := 0
		for k := 0; k < m-1; k += 1 {
			char := worksheet[k][i]
			allSpaces = allSpaces && char == ' '

			if char != ' ' {
				digit, err := strconv.Atoi(string([]byte{char}))
				if err != nil {
					log.Fatal(err)
				}
				operand = operand*10 + digit
			}
		}
		char := worksheet[m-1][i]
		if char != ' ' {
			operator = string([]byte{char})
		}

		if !allSpaces {
			operands = utils.SliceAppend(operands, operand)
			fmt.Printf("%d\n", operand)
		}

		if allSpaces == true || i == n-1 {
			result := operate(operator, operands)
			fmt.Printf("%s %d %d\n\n", operator, operands, result)
			grandTotal += result
			operands = []int{}
			operator = ""
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", grandTotal)
}
