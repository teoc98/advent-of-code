package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"github.com/teoc98/advent-of-code/2025/utils"
)

func isRoll(x byte) bool {
	return x == '@'
}

func neighbourRolls(diagram []string, i, j int) int {
	x := 0
	for ni := max(0, i - 1); ni <= min(len(diagram) - 1, i + 1); ni += 1 {
		for nj := max(0, j - 1); nj <= min(len(diagram[ni]) - 1, j + 1); nj += 1 {
			x += utils.Btoi(isRoll(diagram[ni][nj]))
		}
	}
	return x
}

func isRollAndAccessible(diagram []string, i, j int) bool {
	return isRoll(diagram[i][j]) && neighbourRolls(diagram, i, j) <= 4
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

	accessibleRolls := 0

	var diagram []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		
		diagram = utils.SliceAppend(diagram, line)
	}

	n := len(diagram)
	for i := 0; i < n; i+=1 {
		m := len(diagram[i])
		for j := 0; j < m; j += 1 {
			if isRollAndAccessible(diagram, i, j) {
				accessibleRolls += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", accessibleRolls)
}
