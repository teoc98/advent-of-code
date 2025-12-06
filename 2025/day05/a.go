package main

import (
	"bufio"
	"fmt"
	"github.com/teoc98/advent-of-code/2025/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func isIngredientFresh(freshRanges [][]int, ingredient int) bool {
	for i := 0; i < len(freshRanges); i += 1 {
		if freshRanges[i][0] <= ingredient && ingredient <= freshRanges[i][1] {
			return true
		}
	}
	return false
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

	var freshRanges [][]int
	freshIngredients := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		freshRange, err := utils.MapErr(strings.Split(line, "-"), strconv.Atoi)
		if err != nil {
			log.Fatal(err)
		}
		if len(freshRange) != 2 {
			log.Fatal("invalid range length")
		}
		freshRanges = utils.SliceAppend(freshRanges, freshRange)
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		ingredient, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if isIngredientFresh(freshRanges, ingredient) {
			freshIngredients += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", freshIngredients)
}
