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

func adjcentRolls(line string) []int {
	if line == "" {
		return nil
	}
	n := len(line)
	adj := make([]int, n)
	x := utils.Btoi(isRoll(line[0]))
	for i := 0; i < len(line); i += 1 {
		if i-2 >= 0 {
			x -= utils.Btoi(isRoll(line[i-2]))
		}
		if i+1 < n {
			x += utils.Btoi(isRoll(line[i+1]))
		}
		adj[i] = x
	}
	return adj
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

	var currLine, nextLine string
	var prevAdj, currAdj, nextAdj []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() || currLine != "" {
		currLine = nextLine
		nextLine = scanner.Text()
		prevAdj = currAdj
		currAdj = nextAdj
		nextAdj = adjcentRolls(nextLine)
		
		n := len(currLine)
		for j := 0; j < n; j += 1 {
			if isRoll(currLine[j]) {
				neighbourRolls := 0
				adj := [][]int{prevAdj, currAdj, nextAdj}
				for ni := 0; ni <= 2; ni += 1 {
					if adj[ni] != nil {
						neighbourRolls += adj[ni][j]
					}
				}

				if neighbourRolls <= 4 {
					accessibleRolls += 1
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", accessibleRolls)
}
