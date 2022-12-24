package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

const (
	WALL int = iota
	EMPTY
	UP
	DOWN
	LEFT
	RIGHT
	BLIZZARD
)

var decode = map[string]int{
	"#": WALL,
	".": EMPTY,
	"^": UP,
	"v": DOWN,
	"<": LEFT,
	">": RIGHT,
}

type Grid [][]int

func (g Grid) LoadLine(l string) Grid {
	row := make([]int, len(l))
	for i, r := range l {
		s := string(r)
		row[i] = decode[s]
	}
	g = append(g, row)
	return g
}

func HasElf(elves []*twod.Pos, pos *twod.Pos) bool {
	for _, e := range elves {
		if e.Equals(pos) {
			return true
		}
	}
	return false
}

func (g Grid) Print(bm BlizzardMap, elves map[twod.Pos]bool) {
	s := ""
	for r, row := range g {
		for c, cell := range row {
			add := ""

			if elves[*twod.NewPos(r, c)] {
				add = "E"
			} else {
				switch cell {
				case WALL:
					add = "#"
				case EMPTY:
					add = "."
				case BLIZZARD:
					p := twod.Pos{Row: r, Col: c}
					if b, ok := bm[p]; !ok {
						panic("inconsistent state")
					} else {
						if len(b) > 1 {
							add = fmt.Sprintf("%d", len(b))
						} else {
							add = b[0].String()
						}
					}
				}
			}
			s = fmt.Sprintf("%s%s", s, add)
		}
		s = fmt.Sprintf("%s\n", s)
	}
	fmt.Printf("%v\n", s)
}

func (g Grid) ExtractBlizzard(min, max *twod.Pos) (Grid, BlizzardMap) {
	bm := make(BlizzardMap)
	for r, row := range g {
		for c, cell := range row {
			if cell != WALL && cell != EMPTY {
				b := NewBlizzard(r, c, cell, min, max)
				if _, ok := bm[*b.Pos]; !ok {
					bm[*b.Pos] = make([]*Blizzard, 0)
				}
				bm[*b.Pos] = append(bm[*b.Pos], b)
				g[r][c] = BLIZZARD
			}
		}
	}
	return g, bm
}

func (g Grid) FindTarget(row int) *twod.Pos {
	for c := 0; c < len(g[row]); c++ {
		if g[row][c] == EMPTY {
			return twod.NewPos(row, c)
		}
	}
	panic("no target found")
}

func (g Grid) BlowWind(bm BlizzardMap) (Grid, BlizzardMap) {
	allBliz := make([]*Blizzard, 0)
	for k, v := range bm {
		g[k.Row][k.Col] = EMPTY
		allBliz = append(allBliz, v...)
	}

	newb := make(BlizzardMap, len(bm))
	for _, b := range allBliz {
		b.Move()
		if _, ok := newb[*b.Pos]; !ok {
			newb[*b.Pos] = make([]*Blizzard, 0, 1)
		}
		newb[*b.Pos] = append(newb[*b.Pos], b)
		g[b.Pos.Row][b.Pos.Col] = BLIZZARD
	}
	return g, newb
}

type Blizzard struct {
	Pos *twod.Pos
	Dir *twod.Pos

	ResetAt *twod.Pos
	ResetTo *twod.Pos
	dir     int
}

func (b *Blizzard) Move() {
	b.Pos.Add(b.Dir)
	if b.Pos.Equals(b.ResetAt) {
		b.Pos = b.ResetTo.Clone()
	}
}

func (b *Blizzard) String() string {
	switch b.dir {
	case UP:
		return "^"
	case DOWN:
		return "v"
	case LEFT:
		return "<"
	case RIGHT:
		return ">"
	}
	panic("wrong")
}

func NewBlizzard(r, c int, dir int, min, max *twod.Pos) *Blizzard {
	bliz := Blizzard{Pos: twod.NewPos(r, c), dir: dir}
	switch dir {
	case UP:
		bliz.Dir = twod.NewPos(-1, 0)
		bliz.ResetAt = twod.NewPos(min.Row, c)
		bliz.ResetTo = twod.NewPos(max.Row-2, c)
	case DOWN:
		bliz.Dir = twod.NewPos(1, 0)
		bliz.ResetAt = twod.NewPos(max.Row-1, c)
		bliz.ResetTo = twod.NewPos(min.Row+1, c)
	case LEFT:
		bliz.Dir = twod.NewPos(0, -1)
		bliz.ResetAt = twod.NewPos(r, min.Col)
		bliz.ResetTo = twod.NewPos(r, max.Col-2)
	case RIGHT:
		bliz.Dir = twod.NewPos(0, 1)
		bliz.ResetAt = twod.NewPos(r, max.Col-1)
		bliz.ResetTo = twod.NewPos(r, min.Col+1)
	}
	return &bliz
}

var check = []*twod.Pos{
	twod.NewPos(0, 1), twod.NewPos(-1, 0), twod.NewPos(0, -1), twod.NewPos(1, 0),
}

type BlizzardMap map[twod.Pos][]*Blizzard

func IsValid(p *twod.Pos, start, target, max *twod.Pos) bool {
	return p.Equals(start) || p.Equals(target) ||
		(p.Row > 0 && p.Col > 0 && p.Row < max.Row-1 && p.Col < max.Col-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make(Grid, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = grid.LoadLine(line)
	}
	min := twod.NewPos(0, 0)
	max := twod.NewPos(len(grid), len(grid[0]))
	grid, blizzards := grid.ExtractBlizzard(min, max)

	start := grid.FindTarget(0)
	couldBe := make(map[twod.Pos]bool)
	couldBe[*start] = true
	target := grid.FindTarget(len(grid) - 1)
	log.Printf("Start %v Target %v", start, target)

	grid.Print(blizzards, couldBe)
	grid, blizzards, firstPass := search(grid, blizzards, couldBe, start, target, max)
	log.Printf("Part 1: %v", firstPass)

	// Go back to start
	couldBe = make(map[twod.Pos]bool)
	couldBe[*target] = true
	grid, blizzards, secondPass := search(grid, blizzards, couldBe, target, start, max)

	couldBe = make(map[twod.Pos]bool)
	couldBe[*start] = true
	_, _, thirdPass := search(grid, blizzards, couldBe, start, target, max)

	log.Printf("Part 2: %v", firstPass+secondPass+thirdPass)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// Does a BFS from couldBe (set) to target
// This version uses quantum elves, that can be in all valid positions at the same time :)
func search(grid Grid, blizzards BlizzardMap, couldBe map[twod.Pos]bool, start, target, max *twod.Pos) (Grid, BlizzardMap, int) {
	minute := 0
	for !couldBe[*target] {
		if len(couldBe) == 0 {
			panic("we lost all the elves...")
		}
		minute++
		log.Printf("Minute %v", minute)

		grid, blizzards = grid.BlowWind(blizzards)

		// For all the elves, move to all the new possible positions.
		nextElf := make(map[twod.Pos]bool)
		for e := range couldBe {
			if IsValid(&e, start, target, max) {
				if grid[e.Row][e.Col] == EMPTY {
					nextElf[e] = true
				}
			}
			for _, c := range check {
				pos := e.Clone()
				pos.Add(c)
				if IsValid(pos, start, target, max) {
					if grid[pos.Row][pos.Col] == EMPTY {
						// safe move in this round
						nextElf[*pos] = true
					}
				}
			}
		}
		couldBe = nextElf
		grid.Print(blizzards, couldBe)
	}
	return grid, blizzards, minute
}
