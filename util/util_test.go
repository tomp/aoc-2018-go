package util

import (
	"testing"
)

const (
    TEST_FILE string = "three_lines.txt"
    MISSING_FILE string = "missing.txt"
)

func TestReadLines(t *testing.T) {
    lines, err := ReadLines(TEST_FILE)
    expected := 3
    if err != nil {
        t.Errorf("error reading '%s': %s", TEST_FILE, err)
    } else if len(lines) != expected {
        t.Errorf("read %d lines from '%s' (expected %d)", len(lines),
            TEST_FILE, expected)
    }
}

func TestReadLinesError(t *testing.T) {
    _, err := ReadLines(MISSING_FILE)
    if err == nil {
        t.Errorf("No error reading non-existent file '%s'", MISSING_FILE)
    }
}

func TestParseDigits(t *testing.T) {
    text := "54321"
    digits, err := ParseDigits(text)
    if err != nil {
        t.Errorf("Error parsing '%s': %s", text, err)
    }
    if len(digits) != 5 || digits[0] != 5 || digits[4] != 1 {
        t.Errorf("Misparsed '%s'")
    }
}

func TestParseDigitsError(t *testing.T) {
    text := "5+3.1"
    _, err := ParseDigits(text)
    if err == nil {
        t.Errorf("No error parsing '%s'", text)
    }
}

func TestParseInts(t *testing.T) {
    text := "4 20 7 365 -1"
    vals, err := ParseInts(text)
    if err != nil {
        t.Errorf("Error parsing '%s': %s", text, err)
    }
    if len(vals) != 5 || vals[0] != 4 || vals[4] != -1 {
        t.Errorf("Misparsed '%s'")
    }
}

func TestParseIntsError(t *testing.T) {
    text := "4 score and 7 years ago"
    _, err := ParseInts(text)
    if err == nil {
        t.Errorf("No error parsing '%s'", text)
    }
}

