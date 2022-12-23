package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mikehelmick/AdventOfCode2022/pkg/mathaid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

type Grid map[string]*Elf

func AddRow(g Grid, row int, s string) {
	for c, r := range s {
		if r == '#' {
			p := twod.NewPos(row, c)
			g[p.String()] = NewElf(p)
		}
	}
}

var Adjacent = []*twod.Pos{
	twod.NewPos(-1, -1), twod.NewPos(-1, 0), twod.NewPos(-1, 1),
	twod.NewPos(0, 1), twod.NewPos(1, 1), twod.NewPos(1, 0),
	twod.NewPos(1, -1), twod.NewPos(0, -1),
}

type Check struct {
	IfEmpty []*twod.Pos
	Move    *twod.Pos
}

/*
If there is no Elf in the N, NE, or NW adjacent positions, the Elf proposes moving north one step.
If there is no Elf in the S, SE, or SW adjacent positions, the Elf proposes moving south one step.
If there is no Elf in the W, NW, or SW adjacent positions, the Elf proposes moving west one step.
If there is no Elf in the E, NE, or SE adjacent positions, the Elf proposes moving east one step.
*/
var CheckOrder = []*Check{
	{
		IfEmpty: []*twod.Pos{twod.NewPos(-1, -1), twod.NewPos(-1, 0), twod.NewPos(-1, 1)},
		Move:    twod.NewPos(-1, 0),
	},
	{
		IfEmpty: []*twod.Pos{twod.NewPos(1, -1), twod.NewPos(1, 0), twod.NewPos(1, 1)},
		Move:    twod.NewPos(1, 0),
	},
	{
		IfEmpty: []*twod.Pos{twod.NewPos(0, -1), twod.NewPos(-1, -1), twod.NewPos(1, -1)},
		Move:    twod.NewPos(0, -1),
	},
	{
		IfEmpty: []*twod.Pos{twod.NewPos(0, 1), twod.NewPos(-1, 1), twod.NewPos(1, 1)},
		Move:    twod.NewPos(0, 1),
	},
}

type Elf struct {
	Pos *twod.Pos
}

func NewElf(p *twod.Pos) *Elf {
	return &Elf{
		Pos: p,
	}
}

func (e *Elf) Move(newp *twod.Pos) {
	e.Pos = newp
}

func (g Grid) AllEmpty(p *twod.Pos, check []*twod.Pos) bool {
	for _, c := range check {
		cand := p.Clone()
		cand.Add(c)
		if _, ok := g[cand.String()]; ok {
			return false
		}
	}
	return true
}

func (g Grid) UpdateBounds(tl, br *twod.Pos) {
	for k := range g {
		p := twod.FromString(k)
		tl.Row = mathaid.Min(tl.Row, p.Row)
		tl.Col = mathaid.Min(tl.Col, p.Col)
		br.Row = mathaid.Max(br.Row, p.Row)
		br.Col = mathaid.Max(br.Col, p.Col)
	}
}

func (g Grid) CountEmpty(tl, br *twod.Pos) int {
	count := 0
	for r := tl.Row; r <= br.Row; r++ {
		for col := tl.Col; col <= br.Col; col++ {
			p := twod.NewPos(r, col)
			if _, ok := g[p.String()]; !ok {
				count++
			}
		}
	}
	return count
}

func (g Grid) Print(tl, br *twod.Pos) {
	for r := tl.Row; r <= br.Row; r++ {
		fmt.Printf("%4d ", r)
		for col := tl.Col; col <= br.Col; col++ {
			p := twod.NewPos(r, col)
			if _, ok := g[p.String()]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := make(Grid)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		AddRow(grid, row, line)
		row++
	}

	topLeft := twod.NewPos(100, 100)
	botRight := twod.NewPos(0, 0)
	grid.UpdateBounds(topLeft, botRight)
	grid.Print(topLeft, botRight)

	order := []int{0, 1, 2, 3}
	// 1000 was enough to solve part 2 for my input.
	for i := 0; i < 1000; i++ {
		//log.Printf("Starting round %v, order: %+v", i+1, order)
		moves := make(map[string][]*twod.Pos)
		stable := 0
		// queue up candidate moves
		for spt := range grid {
			p := twod.FromString(spt)
			if grid.AllEmpty(p, Adjacent) {
				stable++
				continue
			}

			cand := p.Clone()
			for _, chidx := range order {
				ch := CheckOrder[chidx]
				if grid.AllEmpty(p, ch.IfEmpty) {
					cand.Add(ch.Move)
					break
				}
			}
			if !p.Equals(cand) {
				if _, ok := moves[cand.String()]; !ok {
					moves[cand.String()] = make([]*twod.Pos, 0)
				}
				moves[cand.String()] = append(moves[cand.String()], p)
			}
		}

		if stable == len(grid) {
			log.Printf("Part 2 : %v", i+1)
			break
		}

		// simple assert that we don't lose anyone.
		before := len(grid)
		//fmt.Printf("round %v movers: %+v\n\n", i, moves)
		for k, movers := range moves {
			if len(movers) == 1 {
				if _, ok := grid[k]; ok {
					panic("invariant violated")
				}
				// Move the elf
				newP := twod.FromString(k)
				grid[k] = grid[movers[0].String()]
				grid[k].Move(newP)
				delete(grid, movers[0].String())
			} // else, just don't move them
		}
		after := len(grid)
		if before != after {
			panic("elf lost")
		}

		// Need to reset these as the bounds may shrink during a round.
		topLeft = twod.NewPos(100, 100)
		botRight = twod.NewPos(0, 0)
		grid.UpdateBounds(topLeft, botRight)
		//fmt.Printf("\n After %v\n", i+1)
		//grid.Print(topLeft, botRight)

		// Rotate the direction check order for the next round
		end := order[0]
		order = append(order[1:], end)

		if i == 9 {
			log.Printf("Bounds: %v %v", topLeft, botRight)
			log.Printf("part 1: %v", grid.CountEmpty(topLeft, botRight))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
