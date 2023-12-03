package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/Xjs/aoc2021/part"
)

func part1(r io.Reader) (int, error) {
	sum := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		first := '\x00'
		last := '\x00'

		for _, char := range s.Text() {
			if unicode.IsDigit(char) {
				if first == 0 {
					first = char
				}
				last = char
			}
		}

		s := string(first) + string(last)
		num, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("error converting %q to int: %w", s, err)
		}
		sum += num
	}

	return sum, nil
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func digitAt(s string, i int) int {
	if unicode.IsDigit(rune(s[i])) {
		return int(s[i] - '0')
	}
	for k, n := range numbers {
		if len(s) < i+len(k) {
			continue
		}
		if s[i:i+len(k)] == k {
			return n
		}
	}

	return -1
}

func part2(r io.Reader) (int, error) {
	sum := 0

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		first := 0
		last := 0
		for i := 0; i < len(line); i++ {
			d := digitAt(line, i)
			if d > 0 {
				if first == 0 {
					first = 10 * d
				}
				last = d
			}
		}
		sum += first + last
	}

	return sum, nil
}

func main() {
	r := os.Stdin
	if part.One() {
		n, err := part1(r)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(n)
		return
	}

	n, err := part2(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(n)
	return
}
