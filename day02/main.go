package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	ROCK int = iota
	PAPER
	SCISSORS
)

var (
	values = map[string]int{
		"A": ROCK, "B": PAPER, "C": SCISSORS,
		"X": ROCK, "Y": PAPER, "Z": SCISSORS,
	}
	shapes = map[int]int{
		ROCK:     1,
		PAPER:    2,
		SCISSORS: 3,
	}

	outcomes = map[string]map[string]string{
		// lose
		"X": map[string]string{"A": "Z", "B": "X", "C": "Y"},
		// draw
		"Y": map[string]string{"A": "X", "B": "Y", "C": "Z"},
		// win
		"Z": map[string]string{"A": "Y", "B": "Z", "C": "X"},
	}
)

func winLoseDraw(opp, you int) int {
	if opp == you {
		return 3
	}
	if (you == ROCK && opp == SCISSORS) ||
		(you == PAPER && opp == ROCK) ||
		(you == SCISSORS && opp == PAPER) {
		return 6
	}
	return 0
}

type Round struct {
	Opponent string
	You      string
}

func New(opp, outcome string) *Round {
	return &Round{
		Opponent: opp,
		You:      outcomes[outcome][opp],
	}
}

func (r *Round) Score() int {
	o := values[r.Opponent]
	y := values[r.You]

	shapeScore := shapes[y]
	return winLoseDraw(o, y) + shapeScore
}

func (r *Round) Print() {
	log.Printf("%v %v %v\n", r.Opponent, r.You, r.Score())
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var score int64
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		r := New(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		score += int64(r.Score())
		r.Print()
	}

	log.Printf("final score: %v\n", score)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
