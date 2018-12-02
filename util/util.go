package util

import (
	"bufio"
	"os"
    "strconv"
    "strings"
)

// ReadLines returns the contents of the given file as a slice
// of lines.
func ReadLines(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err == nil {
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            lines = append(lines, scanner.Text())
        }
        err = scanner.Err()
    }
	return
}

// ParseDigits takes a string consisting of integer digits
// and returns a slice of the integer value for each digit.
func ParseDigits(text string) (digits []int, err error) {
	digits = make([]int, len(text))
	for i, ch := range text {
		val, err := strconv.Atoi(string(ch))
		if err != nil {
			return digits, err
		}
		digits[i] = val
	}
	return digits, nil
}

// ParseInts takes a string consisting of space-separated
// integers and return a slice of the integer values.
func ParseInts(text string) ([]int, error) {
    vals := []int{}
    scanner := bufio.NewScanner(strings.NewReader(text))
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return vals, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}
