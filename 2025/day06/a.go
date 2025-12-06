package main

import (
	"bufio"
	"fmt"
	"github.com/teoc98/advent-of-code/2025/utils"
	"log"
	"os"
	"regexp"
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

	var worksheet [][]int
	var operators []string
	grandTotal := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		tokens := regexp.MustCompile(" +").Split(line, -1)
		if tokens[len(tokens)-1] == "" {
			tokens = tokens[:len(tokens)-1]
		}
		_, err := strconv.Atoi(tokens[0])
		if err == nil {
			operands, err := utils.MapErr(tokens, strconv.Atoi)
			if err != nil {
				log.Fatal(err)
			}
			worksheet = utils.SliceAppend(worksheet, operands)
		} else {
			operators = tokens
		}

		for j := 0; j < len(operators); j += 1 {
			operands := make([]int, len(worksheet))
			for i := 0; i < len(worksheet); i += 1 {
				operands[i] = worksheet[i][j]
			}
			operator := operators[j]
			result := operate(operator, operands)
			grandTotal += result
			fmt.Printf("%d %s %d \n", operands, operators[j], result)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", grandTotal)
}
