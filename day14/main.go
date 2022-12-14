package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

const (
	AIR int = iota
	ROCK
	SAND
	VOID
	SOURCE
)

type Rocks struct {
	Points []*twod.Pos
}

func (r *Rocks) String() string {
	return fmt.Sprintf("%+v", r.Points)
}

func load(l string) *Rocks {
	points := make([]*twod.Pos, 0)

	parts := strings.Split(l, " -> ")
	for _, pt := range parts {
		ptarts := strings.Split(pt, ",")

		c, err := strconv.ParseInt(ptarts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		r, err := strconv.ParseInt(ptarts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		pos := &twod.Pos{
			Row: int(r),
			Col: int(c),
		}
		points = append(points, pos)
	}
	return &Rocks{
		Points: points,
	}
}

func printRow(row []int, min int) {
	for i := min; i < len(row); i++ {
		s := " "
		switch row[i] {
		case AIR:
			s = "."
		case ROCK:
			s = "#"
		case SAND:
			s = "o"
		case VOID:
			s = "@"
		case SOURCE:
			s = "+"
		}

		fmt.Printf("%v", s)
	}
	fmt.Printf("\n")
}

type Grid [][]int

var check = []*twod.Pos{
	{Row: 1, Col: 0},
	{Row: 1, Col: -1},
	{Row: 1, Col: 1},
}

func (g Grid) dropSand(row, col int) *twod.Pos {
	pt := &twod.Pos{Row: row, Col: col}
	for {
		before := pt.Clone()
		for _, c := range check {
			cand := pt.Clone()
			cand.Add(c)
			if cand.Col < 0 || cand.Col >= len(g[0]) {
				panic("out of bounds")
			}
			if g[cand.Row][cand.Col] == AIR {
				pt = cand
				break
			} else if g[cand.Row][cand.Col] == VOID {
				return cand
			}
		}
		if pt.Equals(before) {
			g[pt.Row][pt.Col] = SAND
			return pt
		}
	}
}

func (g Grid) Print(min int) {
	for _, row := range g {
		printRow(row, min)
	}
}

func drawLine(g Grid, r *Rocks) {
	pt := r.Points[0]
	g[pt.Row][pt.Col] = ROCK
	for i := 1; i < len(r.Points); i++ {
		next := r.Points[i]
		for !pt.Equals(next) {
			if pt.Col < next.Col {
				pt.Col++
			} else if pt.Col > next.Col {
				pt.Col--
			}
			if pt.Row < next.Row {
				pt.Row++
			} else if pt.Row > next.Row {
				pt.Row--
			}
			g[pt.Row][pt.Col] = ROCK
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]*Rocks, 0)

	minC := 10000000
	var maxR, maxC int
	for scanner.Scan() {
		line := scanner.Text()
		nl := load(line)
		lines = append(lines, nl)

		for _, p := range nl.Points {
			if p.Row > maxR {
				maxR = p.Row
			}
			if p.Col > maxC {
				maxC = p.Col
			}
			if p.Col < minC {
				minC = p.Col
			}
		}
	}
	maxR += 3
	// For part 1 animations, knock these down to 10
	maxC += 400
	minC -= 400
	//log.Printf("maxR: %v\nmaxC: %v minC: %v\n%+v", maxR, maxC, minC, lines)

	g := make(Grid, maxR)
	for rn := 0; rn < maxR; rn++ {
		g[rn] = make([]int, maxC)
		if rn+1 == maxR {
			for i := range g[rn] {
				g[rn][i] = ROCK
			}
		}
	}
	for _, line := range lines {
		drawLine(g, line)
	}
	g[0][500] = SOURCE
	//g.Print(minC)
	target := &twod.Pos{
		Row: 1,
		Col: 501,
	}

	count := 0
	part1 := 0
	part2 := 0
	part1Done := false
	for {
		count++
		if !part1Done {
			part1 = count
			if res := g.dropSand(0, 500); res.Row == len(g)-2 {
				part1Done = true
			}
		} else {
			// continue to drop for part 2
			part2 = count
			// to animation part one - comment out the if (keep the break)
			if res := g.dropSand(0, 500); res.Equals(target) {
				break
			}
		}
		//log.Printf("grain %v", count)
		//g.Print(minC)

	}
	log.Printf("part 1: %v", part1-1)
	log.Printf("part 2: %v", part2+1)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
