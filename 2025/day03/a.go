package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func maxJoltage(bank string) (int, error) {
	k := 2

	digits := make([]byte, k)
	indexes := make([]int, k)

	for t := 0; t < k; t += 1 {
		var i int

		if t == 0 {
			i = 0
		} else {
			i = indexes[t-1] + 1
		}
		for ; i < len(bank)-(k-1-t); i += 1 {
			c := bank[i]
			if digits[t] == 0 || c > digits[t] {
				digits[t] = c
				indexes[t] = i
			}
		}
	}

	return strconv.Atoi(string(digits))
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

	totalJoltage := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		bank := line
		j, err := maxJoltage(bank)
		if err != nil {
			log.Fatal(err)
		}
		totalJoltage += j
		fmt.Printf("%s %d\n", bank, j)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", totalJoltage)
}
