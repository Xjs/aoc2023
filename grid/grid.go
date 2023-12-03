package grid

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Xjs/aoc2023/parse"
)

type Coordinate uint

// A Point represents a point on a grid.
type Point struct {
	X, Y Coordinate
}

// P is a convenience constructor for Point.
func P(x, y Coordinate) Point {
	return Point{X: x, Y: y}
}

// ulen is a convenience function that returns the length of a slice as Coordinate.
func ulen[T any](s []T) Coordinate {
	return Coordinate(len(s))
}

// A Grid represents a two-dimensional rectangular grid of Ts.
type Grid[T any] struct {
	width, height Coordinate
	values        [][]T
}

// Width returns the grid's width.
func (g Grid[T]) Width() Coordinate {
	return g.width
}

// Height returns the grid's height.
func (g Grid[T]) Height() Coordinate {
	return g.height
}

// NewGrid creates a new zero-filled grid with the given dimensions.
func NewGrid[T any](w, h Coordinate) Grid[T] {
	g := Grid[T]{width: w, height: h, values: make([][]T, h)}
	for i := Coordinate(0); i < h; i++ {
		g.values[i] = make([]T, w)
	}
	return g
}

// GridFrom creates a new Grid from the given values, using the entries of
// the outer slice as rows. It will return an error
// if the rows are not of the same length.
func GridFrom[T any](values [][]T) (Grid[T], error) {
	g := Grid[T]{height: ulen(values), values: values}
	for i, row := range values {
		if i == 0 {
			g.width = ulen(row)
		}
		if ulen(row) != g.width {
			return g, fmt.Errorf("length of row %d is unequal to previous: %d", len(row), g.width)
		}
	}

	return g, nil
}

// ReadIntGrid reads digit lists from r until EOF is encountered,
// and creates a grid from them.
func ReadIntGrid(r io.Reader) (*Grid[int], error) {
	var values [][]int
	s := bufio.NewScanner(r)
	for s.Scan() {
		ds, err := parse.DigitList(s.Text())
		if err != nil {
			return nil, err
		}
		values = append(values, ds)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	g, err := GridFrom(values)
	return &g, err
}

// ReadRuneGrid reads from r until EOF is encountered,
// and creates a grid from the contained runes.
func ReadRuneGrid(r io.Reader) (*Grid[rune], error) {
	var values [][]rune
	s := bufio.NewScanner(r)
	for s.Scan() {
		values = append(values, []rune(s.Text()))
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	g, err := GridFrom(values)
	return &g, err
}

// ErrOutOfBounds is returned by At and Set if an out-of-bounds coordinate is accessed.
var ErrOutOfBounds = errors.New("out of bounds access to grid")

// At returns the value at the given point. It returns ErrOutOfBounds if
// an out-of-bounds point is attempted to be read.
func (g Grid[T]) At(p Point) (T, error) {
	var zero T
	if p.Y >= g.height || p.X >= g.width {
		return zero, ErrOutOfBounds
	}
	return g.values[p.Y][p.X], nil
}

// MustAt is At, but panics instead of returning an error.
func (g Grid[T]) MustAt(p Point) T {
	v, err := g.At(p)
	if err != nil {
		panic(err)
	}
	return v
}

// Environment4 returns a slice of points that represent the 4-environment
// of p, i. e. the points to the left, right, top and bottom. Any points would be
// out of bounds are not returned.
func (g Grid[T]) Environment4(p Point) []Point {
	x, y := p.X, p.Y
	result := make([]Point, 0, 4)
	if x > 0 {
		result = append(result, P(x-1, y))
	}
	if x < g.width-1 {
		result = append(result, P(x+1, y))
	}
	if y > 0 {
		result = append(result, P(x, y-1))
	}
	if y < g.height-1 {
		result = append(result, P(x, y+1))
	}
	return result
}

// Environment8 returns a slice of points that represent the 8-environment
// of p, i. e. the points to the left, right, top and bottom, and all diagonals.
//
//	Any points would be out of bounds are not returned.
func (g Grid[T]) Environment8(p Point) []Point {
	result := make([]Point, 0, 8)
	result = append(result, g.Environment4(p)...)

	x, y := p.X, p.Y
	if x > 0 && y > 0 {
		result = append(result, P(x-1, y-1))
	}
	if x < g.width-1 && y < g.height-1 {
		result = append(result, P(x+1, y+1))
	}
	if x > 0 && y < g.height-1 {
		result = append(result, P(x-1, y+1))
	}
	if x < g.width-1 && y > 0 {
		result = append(result, P(x+1, y-1))
	}
	return result
}

// Set sets the given grid point to the given value. It returns ErrOutOfBounds if
// an out-of-bounds point is attempted to be set.
func (g *Grid[T]) Set(p Point, v T) error {
	if g == nil {
		return errors.New("grid is nil")
	}

	if p.Y >= g.height || p.X >= g.width {
		return ErrOutOfBounds
	}

	g.values[p.Y][p.X] = v
	return nil
}

// MustSet is Set, but panics instead of returning an error.
func (g *Grid[T]) MustSet(p Point, v T) {
	if err := g.Set(p, v); err != nil {
		panic(err)
	}
}

// String creates a multi-line string from an int grid.
func StringIntGrid(g Grid[int]) string {
	var b strings.Builder
	var max int
	for x := Coordinate(0); x < g.width; x++ {
		for y := Coordinate(0); y < g.height; y++ {
			if v := g.MustAt(P(x, y)); v > max {
				max = v
			}
		}
	}

	l := len(fmt.Sprint(max))
	sep := ""
	fill := ' '
	if l > 1 {
		sep = " "
	}

	for y := Coordinate(0); y < g.height; y++ {
		for x := Coordinate(0); x < g.width; x++ {
			v := g.MustAt(P(x, y))
			rep := fmt.Sprint(v)
			for i := 0; i < l-len(rep); i++ {
				b.WriteRune(fill)
			}
			b.WriteString(rep)
			b.WriteString(sep)
		}
		b.WriteRune('\n')
	}

	return b.String()
}

// String creates a multi-line string from an int grid.
func StringCharGrid(g Grid[rune]) string {
	var b strings.Builder

	for y := Coordinate(0); y < g.height; y++ {
		for x := Coordinate(0); x < g.width; x++ {
			v := g.MustAt(P(x, y))
			b.WriteRune(rune(v))
		}
		b.WriteRune('\n')
	}

	return b.String()
}

// Foreach calls f exactly once for each point in g.
func (g *Grid[T]) Foreach(f func(p Point)) {
	for x := Coordinate(0); x < g.Width(); x++ {
		for y := Coordinate(0); y < g.Height(); y++ {
			f(P(x, y))
		}
	}
}
