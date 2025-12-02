package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ScanEntry(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.IndexByte(data, ','); i >= 0 {
		// We have a full comma-terminated entry.
		return i + 1, (data[0:i]), nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated entry.
		return i + 1, (data[0:i]), nil
	}

	// If we're at EOF, we have a final, non-terminated entry. Return it.
	if atEOF {
		return len(data), (data), nil
	}

	// Request more data.
	return 0, nil, nil
}

func intPow(base int, exp int) int {
	x := 1
	for i := 0; i < exp; i += 1 {
		x *= base
	}
	return x
}

func splitInHalf(s string) (string, string) {
	x := len(s) / 2
	return s[:x], s[x:]
}

func parseInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

func isIdInvalid(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for k := 1; k < l; k += 1 {
		if l%k == 0 {
			r := true
			for i := 0; r && i < l; i += k {
				if s[i:i+k] != s[:k] {
					r = false
				}
			}
			if r {
				return true
			}
		}
	}
	return false
}

func sumInvalidIdsInRange(startString string, endString string) int {
	startInt, err := strconv.Atoi(startString)
	if err != nil {
		log.Fatal(err)
	}
	endInt, err := strconv.Atoi(endString)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for x := startInt; x <= endInt; x += 1 {
		if isIdInvalid(x) {
			sum += x
		}
	}
	return sum
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

	invalidIdsSum := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(ScanEntry)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		idRange := strings.Split(line, "-")
		if len(idRange) != 2 {
			log.Fatal("range has unexpected length %d", len(idRange))
		}
		start, end := idRange[0], idRange[1]

		invalidIdsSum += sumInvalidIdsInRange(start, end)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d\n", invalidIdsSum)
}
