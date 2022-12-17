package main

import (
	"fmt"
	"log"

	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

type Glyph struct {
	shape  [][]int
	width  int
	height int
}

var (
	glphs = []*Glyph{
		{
			shape: [][]int{
				{1, 1, 1, 1},
			},
			width:  4,
			height: 1,
		},
		{
			shape: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			width:  3,
			height: 3,
		},
		{
			shape: [][]int{
				{0, 0, 1},
				{0, 0, 1},
				{1, 1, 1},
			},
			width:  3,
			height: 3,
		},
		{
			shape: [][]int{
				{1},
				{1},
				{1},
				{1},
			},
			width:  1,
			height: 4,
		},
		{
			shape: [][]int{
				{1, 1},
				{1, 1},
			},
			width:  2,
			height: 2,
		},
	}

	//input = ">>>><<<<"
	//input = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	input = "><<<<>><<<>><>>><<<>>>><<<<>>>><<>>><>>><>><<<>>><<<>>>><>><>>>><<<>><<<<>><<<><<<<>>><<>>><<><<<>>>><<<><>><<>><<<<>>><>>><<<>>>><<<<>>>><><>>><>>><<<<>>>><<<<>>><<<<><<<>>>><<><<<>>>><<<<>>><>>>><<<>>>><>>><<>>><<<>><<<<>><<>>><<<><<>>><<>>>><<>>><>>><<<>>><<<>>>><<>>><<<>><<<<>>><>>>><<><>>>><>>>><>>>><<>><<<>>>><<<>>><<>>>><<<<>>>><<<><>>>><<<>>><<<<>>><<>>>><<<<>><<<>>>><<><<><<>><<<<>>>><<>>>><<<<>><<<<>>><<>>>><<>>><<<>>>><<<<>><>><<<<>><<<<><<<<>>><<>>><<<<><<<<>>>><>>><<<<>><<>>>><<<><>><<<<>>>><><<<><<<<>>>><<<>>><<<><><<<>>><>>><>>>><<<>>><<><><<>>><<<<><>>><>><>><<>>>><<<<>><<<>><<<>>><<>>><>>>><><<>>><<<<><>><>>>><>>>><>><>><<>>><>>><>>><<<<>><<<<><<<>><<><>><>><<<<>>>><<>><<<>>>><<><<<><<<<>><<<>><<>><>>>><<>>>><<>>>><<>>><<<><<<>><<<>>><<<<>>><>><<>><<<><><<<<><>>><<<<><<<>>><<<>>>><<<><<<>><<<<>>>><<>>><<><><<<<><>><<<><<<<>>><<<<>>>><<><<><<<<>>>><<<>>>><<<<>><<<<>><<<<>>><<<<>>><<>>>><<>>>><<<>>><<>>>><<<><<>>>><><<>>>><>>>><<>><>>>><<<>><<<<>>>><<<>><>>>><<<>><<>>><<<>>>><>>><>>>><<<>><<<>>>><<<<><<<<><<<<>>>><<>><<<><<>>>><<<>>><>>>><><<<><<>><<<><<<>>><>><<<><<<>><>><<<<><<<<><>><<<>>><<<>>><>>><<<<><>>><<<>><<<<>>><><<<>>>><<<><>><<<><<<<>><<<<>>>><><<<<>>><<<<>>>><<<<>><>><<><<<<><<>>><<>>>><<<><<>>>><>><<<><<>><>><<<<>><><<>>>><<>><><<>>><<<>><<><<<>><<><<><>>>><<<<>>><>>><<<>>>><<><<<>>><<<<><<<<>>>><<><>><><<>>><>><><<<>><>>>><>><<<<><<>>><<<<>>><<<>><<<<>>><>>><<<<>>><>>><<<<><<<<>><<<<><<<><<<>>>><<<>>>><<<>>>><<<<>><>><<<<>>><<<<>><<<>><<<>><>>>><<><>>><<<>>>><>><><><>>><<>>><<<>><>><<>>>><<<<>>><<<<>><<><<<<>><>><>>><<<><<>>><<><<><<>>><<>>><<<<>>><<><<<>>>><>>><>>>><<<<><<<>>><<>>><>><<>>>><<<>><<<<>><<<<>><<<>><<<<><<<>><<>>><<<<>>>><<><<<>><><<<<>>><<<<>>>><<<>><<>>><>><<<<>>>><>>><<><<<<>>>><>><<<>>>><<>>>><<<>>><><<<>>>><>>><<><<<><>>><<><<<<>><<>>>><<<<>>><<<<><<<>>>><<<<><<<<>><<<<>><>>>><<><>><<<>>><>><<>><<<<>>><<<>><<><<<<>><<<>>>><<>>>><<<<><><<<<>>>><<>>><<<>>><<<>><<>><<<><>>>><<><<<>><<<>>>><>>><<>><<<><><><<<>>>><>>>><<<<>>><<<><<<<>>><<>>><<<>><<<><<<<>>><<<<>><<<<><<>>>><>><<<<>><<<>>>><<<<>>><<<>><<>><<<>>>><>><>>><>>>><<>>>><<<>><<<<>>>><<<>><<<>>>><<<<>><<<>>><>>><>>><<>><>>><<<<>>>><<><><<><<<<>><>>><>>>><<<>>><<>><><>><><<<<>>>><<<>>>><>>>><>>>><<<<>><<<>><<>><<>>>><<<><>>><<>>><<>>><<<>>><<<<><<>><<><<>><>>>><<<<>>>><>><<>><>>><<<>>><>>>><<>>>><>>>><<<><<<>>>><<<<>><<>>><>><<>>><<<>>>><<<<><>>>><<<<>><<><><<<<>>><>>><<<>><><>><>>><<>>>><<<<><<<>>>><<<>><<>>><<<>><<<>><<<<><><>>>><<>>>><<>>><<<>><<<<><<<><<>>>><>>><<<>><<<<>><>>>><<<<>>>><<<<>>>><>>>><>>>><<<>>><>>><<>><<<>><>><<<<>><<<<>><<<><><<><<<<>><>>><>>><<<>>>><><><<>>>><<<>>><<<<>>><>>>><<<>>><<<><<<>>>><<>>><<<>>>><>>>><<><<><<<<>><<<><<<<>><<<>>>><<><<<<>>><<<<>><<>>>><<<<>><>>>><>><<<>>>><<<<><<<<><><<<><<><<<<><<<><<<<>>>><<<<>>>><<>>><<<<><<><<<<>>><<<<>>><>>>><>><>>><><<<>>><>>>><>><<<>><<<<>>><<>><<><<<>>>><<><<><<><<>><<<>>><<<>>><>>>><>><>><<<<>>>><<>>><<<<>>><<>><<<<><<>>>><>>>><<<<>>>><<<<>><>>><<>>>><<<><<<<>><<<<><<>>>><><<>>><<<<>>>><<>>><<><<>><>>><>>>><<>><<<><<<<><<<<>><<<>><<>>>><>><<<<><<<>><<><<<<><<<>><>>>><<>>><<<<><>>><<<<>>>><<<>><>>>><>><<<<>><>><<<<>>>><>>><><>>><<><>><<<>>>><<<>>><<<<><<<<>><>><<>>>><<<<>><>><<<>><<<<>>>><<>>><<<><<<><<<<>><<<<><<>>>><<<<>>>><<<<>>>><<<>><<<<>><>>><<>>><<<<>>><>>>><<<<><<>><<<>>>><><<<<><<<>>><<<<><<<<>><<<>><<<<>>><<<<>>><<><<><<<>>>><><<<>>>><<>>><<<<>>>><<<<>>>><>>>><<<<>>>><>>><<<>><<<<>>><<><><<<<>>>><<<>>>><>>>><<>>><<<<>>><<<<>><<><<<>><>>>><>>><>><>>><<>>>><<>>><<<<><<<>><<>>><>>>><<<><>><<>>>><>>><<>>><<<>>><<><<<<>>><<>><<<<><<>>><<<><<<>>><<>>><<<<>><><<<>>>><<<<>>><><<>><<<><<<<>><><<<<>><<<>><>><><<<>>><<<<>>><>>><<>><<<>><>><<>><<<><<<>><<>>>><<>>><<><<>>>><<>>>><><>>><<><<<<>><><><<<>><<<>>>><<<>><>>><>>>><><<<>><><>>><>>><<>>><><<<<>>>><<><<<<>>>><<>>>><<<<>>>><<><<<<>>><<<>>><<>>>><<<>>><<<<>><<><<<<><>>>><<<<><<>><>>>><>>><<<<>><<<>>>><<<<>>><<><<<<>>><<<>>><>>><<<<>>><<<<>><>>><<<<>>><>><<<>>><<<<>><<<<>><<<<><>>><>>>><<>>>><<>><<>>>><><<>>><<<>><>><>>><<<<>>><<<>>>><<<>>>><<>>><<<<>>><>>>><<<<>><<>>><><<<>>><>><>><<>><><<<<>>>><<><<>><<<>>>><<<>><>><<><<<<>><<<>><<<>><<<>>>><<<>>><<<<><<><<<<>>>><<<><<>>>><<<><<<>>><<<>>><<<>>><>>>><<<<>>><<<>><<<<>>>><<<><>>>><<><<<<>>>><<<>>><<<>>>><<<<>>><<<<><><<>>>><>><<<<>>>><<><<<<>>>><<<<>>><<<<><>>>><>><>>><<<>>><<<<>><>><>><>>><<>><>>><<>>><><<<<><<<><<>>><<><>>><<><<>>>><>>><<>>>><<<>>><<<>>>><<>>>><<>>>><<>>>><<<<>><<<<><><<<><>>>><<<><>><<<>>><<<<><<><<>>><<>>><<>><<<<>>>><<>>><<>>><<<>>>><<<>>><<<>>>><<<>>><<<<>>>><><<<><<<>><<>>>><<>>><<<<><<<>>>><<>>><<>>><<<>><<>>><<><>>><<<<>>>><<<>>><<<><<<<>>>><<<>>><<><>>><<><<<>><<>>><<<<>><>>>><>>><<<>>><<<>>>><<<<>>><<<>>><>>><>>><<<><<<>><<<><<<<>>><<>>><>>>><<<><<<<>>><<>>><>><<<>>><<><<><<>><<<<>><<<>>><<><<<<><<<<>><<<<>>>><><<<><>>><<>><<<<>>><>><>>>><<><<><<<>><<>><<>>>><<<<><<<>>>><<<>>>><<<>>>><<<><<<<>>>><<<>><<<<><<<>><>>>><>>><<><>>>><>>><<<<>>><<<<>>><<<>><<>>>><<<>>><><>><>><<>><>>>><>>>><>>>><>><<>>>><<<>><<><>><<<<>>><<<>>><<<<>><<<<><<<<>><>>><<<><><<<><<><<>>><<><<><<<>>>><<<>>><<<<>>><><<<>><<<>>>><<<>><>><<<>>><<<<>><>>><<<>><>><<<<>>>><<><>>>><<<><<><<>><>>>><<<<>>>><><<>>>><<>><<>>><><<<>>>><<<>>><<<>>><>><><<<<>><<<<>>><<<<>><<<<><<<>>>><<>>>><<>><<>>>><<><>>>><<>><<<<>>><<<>><>>><<<<>>>><<>><<<<><<<<>>>><<>>><<<>>><<<<>>>><><<<<>><<<><<<<><>>>><<><<><<>><>>>><><>>><<<>>><<<>>>><<>>><<<<>><<<<>>>><<<><<<>><<>>>><>><<><<<>><<<<><>><<>><<<>><<><<<<>>>><<>>>><<<>>>><<<>><<<<><<<<>>>><>>><>>><<<>>><<>><>>><<<>>>><<<>><>>>><><>><<<>><<><<<>><<<>>><<<<>>><<>>><<>>><<<<>><>>><<>>><>>><<<>>><<<<><<>>>><<<<>>><<<<>><<<>>><<>>>><>>><<<<><>>><<<>><<><<><<><><<<>>><<<>>>><<><>><<<>><<<<>><<>>><>>><<<<>><<<><<<<>><>>><>>>><<<<><<<<>>>><<<<>>>><<><<>>>><<<<>><>><<>>><>>>><<>><>><<<>><<<>>><<<<>><>>><<<<>>>><<<>>>><<<<><><<<<><<>>><<<<>>>><<<<>><<<<><<<<><>><><<>><<<>><<<<>>>><<<<>>><<<>>>><<<>><<<>>><>>><<<<>><<<<><<<>><<<<>>>><<<>><<<<><<>><<<>>>><<<<><<<<>>><>>><<<>>><<<>>>><>><>>><<><<><<<<><<>>>><<>>>><<<>>>><<<>>><<<>><<<><<<><><<<>>>><<<<><>><<<>>><>>>><<<>>>><>>>><<>><<><<>>>><<<<>>>><<<<><<><>>><<<<>><<>>><<>><<<>><>><<>>><<<>>>><<<<>><<<>><<<<><<<>><<<<>>>><<<>>>><<<>><><<<>>>><<<>><<<<><><>>><>>><<>><<<>>>><>>>><>>><<<<>>><<<<>>>><<<>><>>>><<>><>>>><<<>>>><<>>>><<<<>>>><<<>><<<<><<<<>>>><<<>>><><<<>>>><>>><<<<>><<<<>>><><<<<>><<<>>><<<<>>><<<><<<<>>><<<>>>><<<><<<><><<<<>>><<<<>>><<<>><<<<>><><<<>>><>>><<<>>><<<>>>><>>>><<<><<>>><<>>><>>>><<<>><>><><<<<>>><<<>>>><<>>><<><>>>><<<>><<<><<><<<>>>><>>>><<<<>><<<<>><<<>>>><<<><<<><>>>><<<<>><<><<>>><<<<>>><>>><<<><><<<<>><>><>>><<<>>>><<<><><<<<><<<<>>><<><<<<>><<>>>><<><>>>><<<>>><<<>>>><<<>>><><<<<><<>>><<><<<><<<<>>><<>>>><>><<<>>>><<>>>><<><<<>>><<>>>><>><<<>>>><<<<>>>><>>><<><><<><><>><<>><>><>><<>>>><<>><<<<><<<<>>>><><<>><>>>><<><<><<>>><<<><<<>>>><<<>>>><<<<>>><<<>>><<<<>><<<>>>><<<><>>><><<<<><<<<>>>><<>>>><>>>><<>>>><>>><<>><<<<><<<<>>>><<<<><>>>><<<>><<>>><<<<><<>><<>><><<<>>><<<>>>><<<>>><>>><><<<<>>>><<>>>><>>>><><<<<>>>><<>>>><<>><>>>><<<<><<<>>>><<>><<<<>>>><<<><<<><<<>>><<>>><>><<>>><>><<<>>><<<<>><<>><<<>>>><<<>>>><<><<>>>><<>><<<<>><<<>>>><<>>>><<<>>><<><<>>>><<>>><>>>><<<><<<>><>><<<<>>><<>><>><<>><>>>><><<<>><<><<>>><<>><<<><<><<><<<><><<>>>><<<><<<>>><>>>><<<>>>><<<<>>><<<><>>><>>>><<>><<<<>><>><<<>><<>><<<><><<>>>><<<<>>><<<>>><<<<>><<>><<<>><>><<<>>><<<>>><<<<>>>><<>>>><<<<>>><<<>><<<<><><<<<>>>><>>>><<<<>><<<<><<<<>><>>>><<<><>><<<>>><<<<><<<<>>><<<<><<>>>><<><>><<<>>>><<<<>>>><>>><<<>><<<<>>><>>><>><<<>>><<<>>>><>>>><<<<>><<>><<<>>><<<<>>><<<<><<<>>><<<><<<<><><<<><<><<<><<<<>>><<<<>><<<<>><>><>><<><<<>>><<<><<><<>>><<<<>>><<<>>>><<><<<>>>><<<>><<<>><>>>><>>><>><<<>><<>>>><<<<>>>><>><<>><<<>>><>>>><<>><<<>><<><<>>>><><<<><>>><<>>>><<<<>>>><<<>>>><<<>>><<<><<<>>>><<<>>>><<<>>>><<<<><>>><>><<>><<<><<<><<>><>>><<<>>><<<>><<<>>><<<<>>>><<<<><>>><<>>><>>><<<<>><<>><>>>><<>>><<<<>>><<<<>>><<<<><<<<>>><<<>><><><<<><<<<>>><>>><<><>><<>>>><<>>><<<>>>><<<<>>><<<>><<>><>>><>>><<<<>>>><<<<>>>><<>><<<<>>>><>>>><<<><<<>>><<<>>>><<<<>>><><<<<>><<><<>>>><<<<>>><<<<>>><<<<><<<>>>><<><<>>>><><<<>>>><<><<>>><<>>>><<<><<><>>><<<<>>><<<>>>><<><<>>><<<<><>>>><<>>><<<>>>><<<<><<>><<>><<>>><><<<>><<<<>><><<<<><>><<>><<<>>>><<>>>><>>>><<><>><>>><<<>>><>>>><<<><<><<<>><<<>><<>>><<><<>>><<<>>><<<>>>><>>><>><><<><<>>><<<<>>><>>><<<><<><<<>><>><<>>><>>>><<>>>><<<><>>>><<>>>><>>>><<<<>>><<>>><>>><<>><<><<<<>>>><>>><<>>>><><<><<<>>><<>><>>><<<<>>>><>>>><<<>>><<<>><<<>><<>>>><>><>>><<<><<<>>>><<<<>><<>>><<<<>>><<>>><<>>><>>>><<<<><<<>>>><<<>>><<<<><<><<>>>><<<>>><<<<>>><><<<<>><<<><<<<>><<<>><<<<><<<<>>>><>><>>><<><<<>><<<>>><><<<<>><<<<>><<>>><<<<>>><>>><>><<<>><<<>><>>><<<>>>><<<>><<<<><<><<<<>>><<<>>>><<<<>>><<<<>>><><<<>>><<<>><<<<><<<<><<<>>><<<>>>><><<<>><>>>><<<<>><<>>><<<>>><<>>>><<<<>>>><<>>>><<<<>>><<>>><<>><>>><<>>><<<<>><>>><><>><><<>>>><<>>><>>><<<<><<<<>><<<<>>><<<<>>><<<<>>>><<<<>>><<<><><<<<><>>>><>><<<<>><<<<><<>>>><><>>>><<<<>><<>>>><<<><>>>><<<>><<<>><>><>>>><<<>>><<<>>><<>>>><<<><<<>><>>>><>>><<<>>>><><<>>><<<>>>><<>>>><<<>>>><<<<><<><>><>>>><<>>>><<<<>><<<><<<<>>><>>><<<>>><>><>>>><<<<>>>><<<>><<>>><>>><<<>><<<<>>><<<>>><<>>><<>>><>>>><<>>>><<<<>>>><<>>>><<><<>><>>>><<>>>><<<<>>>><<<<><<<><>>>><<<>>>><<<<>>>><><<<><<<<><<<<>><<<><>>><<>>><>><<>>>><<>><<<<>>>><<>>><<>>>><<>><>><><>><<<<><>>><<>>>><<>>>><>>><><<<><<<<>>><<<>>>><<>>><<<<>><<>>><<<<>>><<<<>><<>>><><<>>><<>><<<>>><<>>><<<>>><<>><<><>>>><>><>>><<>><<>>><<<<>>>><<<>>>><<<<>><<<<>>>><<<<><<<><>>><<<<><<>>><<>><<<<><<<><<<><<><<><>>>><<>><<<><>><<>><<<><<<><<<>>><<<><<>><<<><<<>>>><<>>>><<><>>>><<<<>>><><<><<<>><>>>><<<>>><<<>>>><>><<<>>>><<>>>><><<<>>><<>><><<>><>>><<>>><<><<<>><<<<>><><<<<>>>><<><<<<>>>><<<<>><<<><<<<>><<<<>>><<><<><>><><<<>>>><<<>>>><<<<>><<<>>><<>>><<>>>><>>>><<>>><>>>><<<<>><<<>>><<<><>>><>>><<<><<<<>>>><<<>>>><<<<><>>><<<>><<>><<>>><<>>>><>>><<<>>><<<>>>><<<>>>><<<<>><><<<><>><>><<<<>>>><<<>>>><<><>><><<<<>>>><><>>>><<>><<<<>><<<<><<<<>>>><<<>>><<<>>><<<>><<<<>><<>><>>><<>><<>><<<<>>>><<<>><>>>><<<<>>>><<<>>>><<>>><<<>>>><<<<><<<>><<<>>><<<<>>><><<>>>><<>>><<>>>"
)

func initRow() []int {
	return make([]int, 7)
}

func growChamber(chamber [][]int, height int) [][]int {
	blanks := 0
	done := false
	for _, row := range chamber {
		for _, cell := range row {
			if cell != 0 {
				done = true
				break
			}
		}
		if done {
			break
		}
		blanks++
	}

	need := 3 + height
	need -= blanks

	if need < 0 {
		// need is negative here
		newChamber := make([][]int, len(chamber)+need)
		copy(newChamber, chamber[(need*-1):])
		//print(newChamber)
		return newChamber
	}

	// add blanks + height at the top.
	newChamber := make([][]int, len(chamber)+need)
	for i := 0; i < need; i++ {
		newChamber[i] = initRow()
	}
	copy(newChamber[need:], chamber)
	return newChamber
}

type Reset struct {
	r int
	c int
	v int
}

func NewReset(r, c, v int) *Reset {
	return &Reset{r: r, c: c, v: v}
}

func reset(chamber [][]int, r []*Reset) {
	for i := len(r) - 1; i >= 0; i-- {
		chamber[r[i].r][r[i].c] = r[i].v
	}
}

func placeGlyph(chamber [][]int, g *Glyph, p *twod.Pos) bool {
	resets := make([]*Reset, 0, 10)
	ro := p.Row
	co := p.Col
	needsReset := false
	for r, row := range g.shape {
		for c, v := range row {
			if g.shape[r][c] == 0 {
				continue
			}
			if co+c >= 7 || co+c < 0 || r+ro >= len(chamber) {
				needsReset = true
				break
			}
			if chamber[r+ro][co+c] != 0 {
				needsReset = true
				break
			}
			resets = append(resets, NewReset(r+ro, co+c, chamber[r+ro][co+c]))
			chamber[r+ro][co+c] = v
		}
		if needsReset {
			break
		}
	}
	if needsReset {
		reset(chamber, resets)
		return false
	}
	return true
}

func eraseGlyph(cha [][]int, g *Glyph, p *twod.Pos, rv int) {
	for r := p.Row; r < p.Row+g.height; r++ {
		for c := p.Col; c < p.Col+g.width; c++ {
			if cha[r][c] == 1 {
				cha[r][c] = rv
			}
		}
	}
}

func print(chamber [][]int) {
	for i, r := range chamber {
		fmt.Printf("%6d ", len(chamber)-i)
		fmt.Print("|")
		for _, v := range r {
			if v == 0 {
				fmt.Print(".")
			} else if v == 1 {
				fmt.Print("@")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Print("|")
		fmt.Print("\n")
	}
	fmt.Printf("       |-------|\n\n")
}

const DOWN = "D"

var moves = map[string]*twod.Pos{
	"<":  twod.NewPos(0, -1),
	">":  twod.NewPos(0, 1),
	DOWN: twod.NewPos(1, 0),
}

func TopRows(chamber [][]int, c int) string {
	start := 0
	for ; isEmpty(chamber[start]); start++ {
	}

	s := ""
	for i := start; i < start+c; i++ {
		for _, v := range chamber[i] {
			s = fmt.Sprintf("%s%v", s, v)
		}
	}
	return s
}

func NewState(i int, jet int, chamber [][]int) string {
	return fmt.Sprintf("%1d:%06d:%s", i%len(glphs), jet, TopRows(chamber, 30))
}

func main() {
	chamber := make([][]int, 3)
	for i := range chamber {
		chamber[i] = initRow()
	}

	// These values are calculated by uncommenting the state calculation below and
	// 5500 rounds.
	initialHeight := int64(3018)
	cycleStart := int64(1933)
	cycleLength := int64(1740)
	cycleHeight := int64(2724)

	height := initialHeight
	target := int64(1000000000000)
	target -= cycleStart
	cycles := target / cycleLength
	height += cycles * cycleHeight
	target -= cycles * cycleLength

	extra := int(target - 1) // 1 idx, not 0 idx.
	log.Printf("simulating %v runs", extra)

	//states := make(map[string]int)

	jetI := 0
	for i := 0; i < extra; i++ {
		// For calculating the cycles in part 2.
		/*
			if i == 193 || i == 1933 || i == 3673 || i == 5413 {
				state := NewState(i, jetI, chamber)
				if p, ok := states[state]; ok {
					log.Printf("CYCLE: %v -> %v : len=%v, height=%v", i, p, i-p, RockHeight(chamber))
				}
				states[state] = i
			}
		*/

		g := glphs[i%len(glphs)]
		p := twod.NewPos(0, 2)

		chamber = growChamber(chamber, g.height)
		placeGlyph(chamber, g, p)
		//print(chamber)

		falling := true
		for falling {
			eraseGlyph(chamber, g, p, 0)

			// apply jet.
			dir := string(input[jetI])
			jetI = (jetI + 1) % len(input)
			np := p.Clone()
			np.Add(moves[dir])
			eraseGlyph(chamber, g, p, 0)
			if placeGlyph(chamber, g, np) {
				p = np
			} else {
				// couldn't move, put it back
				placeGlyph(chamber, g, p)
			}
			//print(chamber)

			// apply gravity.
			np = p.Clone()
			np.Add(moves[DOWN])
			eraseGlyph(chamber, g, p, 0)
			falling = placeGlyph(chamber, g, np)
			if falling {
				p = np
			} else {
				// couldn't move, put it back
				placeGlyph(chamber, g, p)
				// harden the position.
				eraseGlyph(chamber, g, p, 2)
			}
			//print(chamber)
		}
		//print(chamber)
	}
	//print(chamber)

	part2 := int64(RockHeight(chamber))
	log.Printf("part 2: %v", part2+height)
}

func RockHeight(chamber [][]int) int {
	empty := 0
	for _, r := range chamber {
		if isEmpty(r) {
			empty++
		} else {
			break
		}
	}
	return len(chamber) - empty
}

func isEmpty(r []int) bool {
	for _, v := range r {
		if v != 0 {
			return false
		}
	}
	return true
}
