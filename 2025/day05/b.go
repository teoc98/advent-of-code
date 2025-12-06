package main

import (
	"bufio"
	"fmt"
	"github.com/teoc98/advent-of-code/2025/utils"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
	totalFreshIngredients := 0

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

	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i][0] < freshRanges[j][0] || (freshRanges[i][0] == freshRanges[j][0] && freshRanges[i][1] < freshRanges[j][1])
	})

	currRange := []int{-1, -2}
	for i := 0; i < len(freshRanges); i += 1 {
		freshRange := freshRanges[i]

		if freshRange[0] <= currRange[1]+1 {
			// extend currRange
			currRange[1] = max(freshRange[1], currRange[1])
		} else {
			// reset currRange
			totalFreshIngredients += currRange[1] - currRange[0] + 1
			currRange = freshRange
		}

	}
	totalFreshIngredients += currRange[1] - currRange[0] + 1

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", totalFreshIngredients)
}
