package main

import (
	"bufio"
	"log"
	"os"

	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

type Grid [][]int

func isValid(g Grid, p *twod.Pos) bool {
	return p.Row >= 0 && p.Col >= 0 &&
		p.Row <= len(g)-1 && p.Col <= len(g[0])-1
}

func makeFun(g Grid) twod.ValidFunc {
	return func(p *twod.Pos) bool {
		return isValid(g, p)
	}
}

func (g Grid) Val(p *twod.Pos) int {
	return g[p.Row][p.Col]
}

// BFS does a multi-origin BFS towards a specific end, e.
func BFS(g Grid, initial []*twod.Pos, e *twod.Pos) int {
	validF := makeFun(g)
	wave := 0

	queue := initial
	visited := make(map[string]bool)
	for _, s := range queue {
		visited[s.String()] = true
	}

	for len(queue) > 0 {
		wave++
		nextWave := make([]*twod.Pos, 0)
		// For each item in this wavefront.
		for _, p := range queue {
			curVal := g.Val(p)
			// Check all valid neighbors.
			cand := p.Neighbors(validF)
			for _, n := range cand {
				candVal := g.Val(n)
				// See if it's a valid step.
				if candVal <= curVal+1 {
					// See if we would step to the target.
					if n.Equals(e) {
						return wave
					}
					// if we haven't already been there, queue
					if !visited[n.String()] {
						nextWave = append(nextWave, n)
						visited[n.String()] = true
					}
				}
			}
		}
		queue = nextWave
	}
	// didn't reach the end.
	return -1
}

func main() {

	g := make(Grid, 0)

	var start, end *twod.Pos

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%v", line)

		row := make([]int, 0)
		for _, c := range line {
			v := int(c - 'a')
			if c == 'S' {
				start = &twod.Pos{
					Row: len(g),
					Col: len(row),
				}
				v = 0
			} else if c == 'E' {
				end = &twod.Pos{
					Row: len(g),
					Col: len(row),
				}
				v = 26
			}
			row = append(row, v)
		}
		g = append(g, row)
	}
	log.Printf("%+v", g)
	log.Printf("%+v", *start)
	log.Printf("%+v", *end)

	part1 := BFS(g, []*twod.Pos{start}, end)
	log.Printf("part 1 %+v", part1)

	initial := make([]*twod.Pos, 0)
	for r, row := range g {
		for c := range row {
			if p := twod.NewPos(r, c); g.Val(p) == 0 {
				initial = append(initial, p)
			}
		}
	}
	part2 := BFS(g, initial, end)
	log.Printf("part 2 %+v", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
