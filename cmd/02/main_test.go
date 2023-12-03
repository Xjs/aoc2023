package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			"sample1",
			`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
			8,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimum(t *testing.T) {
	tests := []struct {
		name string
		g    Game
		want Sample
	}{
		{"1", mustParseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"), Sample{4, 2, 6}},
		{"1", mustParseGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"), Sample{1, 3, 4}},
		{"1", mustParseGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"), Sample{20, 13, 6}},
		{"1", mustParseGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"), Sample{14, 3, 15}},
		{"1", mustParseGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"), Sample{6, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimum(tt.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
