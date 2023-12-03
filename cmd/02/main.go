package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Xjs/aoc2021/part"
)

type Sample struct {
	Red, Green, Blue int
}

type Game struct {
	ID      int
	Samples []Sample
}

func parse(r io.Reader) ([]Game, error) {
	var games []Game

	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		game, err := parseGame(line)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func mustParseGame(line string) Game {
	g, _ := parseGame(line)
	return g
}

func parseGame(line string) (Game, error) {
	splices := strings.Split(line, ":")
	if len(splices) != 2 {
		return Game{}, fmt.Errorf("invalid number of colons in line %q", line)
	}
	gameID, rest := splices[0], splices[1]

	const gamePrefix = "Game "
	if !strings.HasPrefix(gameID, gamePrefix) {
		return Game{}, fmt.Errorf("line %q doesn't start with %q", line, gamePrefix)
	}
	gameID = strings.TrimPrefix(gameID, gamePrefix)

	var game Game
	var err error
	game.ID, err = strconv.Atoi(gameID)
	if err != nil {
		return Game{}, fmt.Errorf("line %q: game ID: %w", line, err)
	}

	for _, splice := range strings.Split(rest, ";") {
		var sample Sample
		for _, colour := range strings.Split(splice, ",") {
			colour = strings.TrimSpace(colour)
			switch {
			case strings.HasSuffix(colour, " red"):
				red, err := strconv.Atoi(strings.TrimSuffix(colour, " red"))
				if err != nil {
					return Game{}, fmt.Errorf("line %q: red: %w", line, err)
				}
				sample.Red = red
			case strings.HasSuffix(colour, " blue"):
				blue, err := strconv.Atoi(strings.TrimSuffix(colour, " blue"))
				if err != nil {
					return Game{}, fmt.Errorf("line %q: blue: %w", line, err)
				}
				sample.Blue = blue
			case strings.HasSuffix(colour, " green"):
				green, err := strconv.Atoi(strings.TrimSuffix(colour, " green"))
				if err != nil {
					return Game{}, fmt.Errorf("line %q: green: %w", line, err)
				}
				sample.Green = green
			default:
				return Game{}, fmt.Errorf("invalid colour in %q", line)
			}
		}
		game.Samples = append(game.Samples, sample)
	}
	return game, nil
}

func possible(g Game) bool {
	var reference = Sample{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for _, sample := range g.Samples {
		switch {
		case sample.Blue > reference.Blue,
			sample.Red > reference.Red,
			sample.Green > reference.Green:
			return false
		}
	}

	return true
}

func sumPossible(g []Game) int {
	sum := 0

	for _, game := range g {
		if possible(game) {
			sum += game.ID
		}
	}

	return sum
}

func part1(r io.Reader) (int, error) {
	games, err := parse(r)
	if err != nil {
		return 0, err
	}
	return sumPossible(games), nil
}

func minimum(g Game) Sample {
	var min Sample
	for _, s := range g.Samples {
		if s.Blue > min.Blue {
			min.Blue = s.Blue
		}
		if s.Red > min.Red {
			min.Red = s.Red
		}
		if s.Green > min.Green {
			min.Green = s.Green
		}
	}
	return min
}

func power(s Sample) int {
	return s.Red * s.Blue * s.Green
}

func part2(r io.Reader) (int, error) {
	games, err := parse(r)
	if err != nil {
		return 0, err
	}
	sumPowers := 0
	for _, game := range games {
		sumPowers += power(minimum(game))
	}
	return sumPowers, nil
}

func main() {
	r := os.Stdin

	if part.One() {
		n, err := part1(r)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("part1: %d", n)
	}

	n, err := part2(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("part2: %d", n)
}
