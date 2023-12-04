package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Xjs/aoc2023/parse"
)

type Card struct {
	ID      int
	Winning map[int]struct{}
	Have    map[int]struct{}

	Copies int
}

func ParseCard(l string) (Card, error) {
	const prefix = "Card "
	if !strings.HasPrefix(l, prefix) {
		return Card{}, fmt.Errorf("line %q doesn't start with %q", l, prefix)
	}
	l = strings.TrimPrefix(l, prefix)

	parts := strings.Split(l, ":")
	if len(parts) != 2 {
		return Card{}, fmt.Errorf("line %q doesn't contain exactly one colon", l)
	}

	id, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Card{}, fmt.Errorf("line %q: error parsing ID: %w", l, err)
	}

	card := Card{ID: id, Copies: 1}

	parts = strings.Split(parts[1], "|")
	if len(parts) != 2 {
		return card, fmt.Errorf("line %q: must contain exactly one pipe", l)
	}

	winning, err := parse.IntListWhitespace(parts[0])
	if err != nil {
		return card, fmt.Errorf("part %q: error parsing integer list: %w", parts[0], err)
	}

	card.Winning = make(map[int]struct{})
	for _, num := range winning {
		card.Winning[num] = struct{}{}
	}

	have, err := parse.IntListWhitespace(parts[1])
	if err != nil {
		return card, fmt.Errorf("part %q: error parsing integer list: %w", parts[1], err)
	}

	card.Have = make(map[int]struct{})
	for _, num := range have {
		card.Have[num] = struct{}{}
	}

	return card, nil
}

func (c Card) Value() int {
	return int(math.Pow(2, float64(c.Matching()-1)))
}

func (c Card) Matching() int {
	matching := 0

	for num := range c.Winning {
		_, have := c.Have[num]
		if !have {
			continue
		}
		matching++
	}
	return matching
}

func parseAll(r io.Reader) ([]*Card, error) {
	var cards []*Card
	s := bufio.NewScanner(r)
	for s.Scan() {
		card, err := ParseCard(s.Text())
		if err != nil {
			return cards, err
		}
		cards = append(cards, &card)
	}
	return cards, nil
}

func main() {
	sum := 0
	cards, err := parseAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for _, card := range cards {
		sum += card.Value()
	}

	for i, card := range cards {
		n := card.Matching()
		for j := i + 1; j < min(i+n+1, len(cards)); j++ {
			cards[j].Copies += card.Copies
		}
	}

	sum2 := 0
	for _, card := range cards {
		sum2 += card.Copies
	}

	log.Printf("part1: %d, part2: %d", sum, sum2)
}
