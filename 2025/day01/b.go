package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	mod := 100
	position := 50
	timesPointedAtZero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dirChar := line[0]
		var dir int
		switch dirChar {
		case 'L':
			dir = -1
		case 'R':
			dir = +1
		default:
			log.Fatalf("unexpected direction: '%c'")
		}

		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		prevPosition := position
		position = position + dir*num

		if position == 0 {
			timesPointedAtZero += 1
		}
		
		if position < 0 {
			if prevPosition == 0 {
				timesPointedAtZero -= 1
			}
			for position < 0 {
				position += mod
				timesPointedAtZero += 1
			}
			if position == 0 {
				timesPointedAtZero += 1
			}
		}

		for position >= mod {
			position -= mod
			timesPointedAtZero += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", timesPointedAtZero)
}
