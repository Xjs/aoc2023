package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"

	"github.com/Xjs/aoc2023/grid"
)

func part1and2(r io.Reader) (int, int, error) {
	type id int

	g, err := grid.ReadRuneGrid(r)
	if err != nil {
		return 0, 0, err
	}

	classG := grid.NewGrid[id](g.Width(), g.Height())
	current := id(1)
	numbers := make(map[id]int)

	for y := grid.Coordinate(0); y < g.Height(); y++ {
		var currentRunes []rune
		// go 1 beyond the width to avoid duplicating the parsing and classification code
		for x := grid.Coordinate(0); x < g.Width()+1; x++ {
			var p grid.Point
			var r rune
			// Only set p and r if we're not out of bounds
			if x != g.Width() {
				p = grid.P(x, y)
				r = g.MustAt(p)
			}

			if unicode.IsDigit(r) {
				currentRunes = append(currentRunes, r)
				classG.Set(p, current)
			} else if len(currentRunes) > 0 {
				n, err := strconv.Atoi(string(currentRunes))
				if err != nil {
					return 0, 0, fmt.Errorf("Error parsing number %q: %w", string(currentRunes), err)
				}
				numbers[current] = n
				current++
				currentRunes = nil
			}
		}
	}

	haveAdjacentSymbol := make(map[id][]grid.Point)

	g.Foreach(func(p grid.Point) {
		theID := classG.MustAt(p)
		if theID == 0 {
			return
		}

		neighbours := g.Environment8(p)
		for _, neighbour := range neighbours {
			r := g.MustAt(neighbour)
			if unicode.IsDigit(r) {
				continue
			}
			if r == '.' {
				continue
			}
			haveAdjacentSymbol[theID] = append(haveAdjacentSymbol[theID], neighbour)
		}
	})

	gearMap := make(map[grid.Point]map[id]struct{})

	sum := 0
	for theID, syms := range haveAdjacentSymbol {
		if len(syms) == 0 {
			continue
		}
		sum += numbers[theID]

		for _, sym := range syms {
			if g.MustAt(sym) == '*' {
				if gearMap[sym] == nil {
					gearMap[sym] = make(map[id]struct{})
				}
				gearMap[sym][theID] = struct{}{}
			}
		}
	}

	gearSum := 0

	for _, ids := range gearMap {
		if len(ids) != 2 {
			continue
		}
		n := 1
		for id := range ids {
			n *= numbers[id]
		}
		gearSum += n
	}

	return sum, gearSum, nil
}

func main() {
	p1, p2, err := part1and2(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("part1: %d, part2: %d", p1, p2)
}
