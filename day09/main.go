package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

type Grid [][]bool

func TooFar(p1, p2 *twod.Pos) bool {
	return math.Abs(float64(p1.Row-p2.Row)) > 1 ||
		math.Abs(float64(p1.Col-p2.Col)) > 1
}

const (
	HEIGHT = 500
	WIDTH  = 500
)

func move(dir string, segments []*twod.Pos, g Grid) {
	segments[0].Add(twod.Dirs[dir])
	for i := 1; i < len(segments); i++ {
		cur := segments[i]
		last := segments[i-1]
		if TooFar(cur, last) {
			// Very literal reading of the movement rules.
			// if same row or column, move closer.
			if cur.Row == last.Row {
				if cur.Col > last.Col {
					cur.Col--
				} else {
					cur.Col++
				}
			} else if cur.Col == last.Col {
				if cur.Row > last.Row {
					cur.Row--
				} else {
					cur.Row++
				}
			} else {
				// otherwise diagonal
				for _, d := range twod.Diags {
					cand := cur.Clone()
					cand.Add(d)
					if !TooFar(last, cand) {
						segments[i] = cand
						break
					}
				}
			}
		}
	}
	last := segments[len(segments)-1]
	g[last.Row][last.Col] = true
}

func printVisited(part string, g Grid) {
	visited := 0
	for _, r := range g {
		for _, c := range r {
			if c {
				visited++
			}
		}
	}
	log.Printf("part %v %v", part, visited)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// would have been nice if the problem said infinite field...
	g := make(Grid, HEIGHT)
	g2 := make(Grid, HEIGHT)
	for i, _ := range g {
		g[i] = make([]bool, WIDTH)
		g2[i] = make([]bool, WIDTH)
	}

	p1segments := []*twod.Pos{
		{Row: HEIGHT / 2, Col: WIDTH / 2},
		{Row: HEIGHT / 2, Col: WIDTH / 2},
	}
	g[p1segments[0].Row][p1segments[0].Col] = true

	segments := make([]*twod.Pos, 10)
	for i, _ := range segments {
		segments[i] = &twod.Pos{Row: HEIGHT / 2, Col: WIDTH / 2}
	}
	// ok to take the wrong one here...
	g2[segments[0].Row][segments[0].Col] = true

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		steps, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			panic(err)
		}

		//log.Printf("%v", line)
		for i := 0; int64(i) < steps; i++ {
			move(parts[0], p1segments, g)
			move(parts[0], segments, g2)

			/* debugging...
			visual := make(map[string]int)
			for i, p := range segments {
				visual[fmt.Sprintf("%v,%v", p.Row, p.Col)] = i
			}
			for r, _ := range g2 {
				for c, v := range g2[r] {
					if k, ok := visual[fmt.Sprintf("%v,%v", r, c)]; ok {
						fmt.Printf("%v", k)
					} else if v {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Println("")
			}
			fmt.Println("----")
			*/
		}
	}

	printVisited("1", g)
	printVisited("2", g2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
