package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Tree struct {
	Height  int64
	Visible bool
}

func (t *Tree) String() string {
	v := "_"
	if t.Visible {
		v = "*"
	}
	return fmt.Sprintf("%v%v", t.Height, v)
}

type Pos struct {
	Row int
	Col int
}

type Grid [][]*Tree

func (g Grid) MarkVisible() {
	// brute force ... for each row, from each side.
	for row := 0; row < len(g); row++ {
		g[row][0].Visible = true
		highest := g[row][0].Height
		for col := 1; col < len(g[row]); col++ {
			if h := g[row][col].Height; h > highest {
				g[row][col].Visible = true
				highest = h
			}
		}
		g[row][len(g[row])-1].Visible = true
		highest = g[row][len(g[row])-1].Height
		for col := len(g[row]) - 2; col > 0; col-- {
			if h := g[row][col].Height; h > highest {
				g[row][col].Visible = true
				highest = h
			}
		}
	}
	// for each column, from each end
	for col := 0; col < len(g[0]); col++ {
		g[0][col].Visible = true
		highest := g[0][col].Height
		for row := 1; row < len(g); row++ {
			if h := g[row][col].Height; h > highest {
				g[row][col].Visible = true
				highest = h
			}
		}
		g[len(g)-1][col].Visible = true
		highest = g[len(g)-1][col].Height
		for row := len(g) - 2; row > 0; row-- {
			if h := g[row][col].Height; h > highest {
				g[row][col].Visible = true
				highest = h
			}
		}
	}
}

func (g Grid) ScenicScore() int {
	best := 0

	for r := 1; r < len(g)-1; r++ {
		for c := 1; c < len(g[r])-1; c++ {
			h := g[r][c].Height
			// again brute force - fastest typing time...
			// for each position, look in every direction.
			above := 0
			for rm := r - 1; rm >= 0; rm-- {
				if g[rm][c].Height < h {
					above++
				} else {
					above++
					break
				}
			}
			below := 0
			for rm := r + 1; rm < len(g); rm++ {
				if g[rm][c].Height < h {
					below++
				} else {
					below++
					break
				}
			}

			left := 0
			for cm := c - 1; cm >= 0; cm-- {
				if g[r][cm].Height < h {
					left++
				} else {
					left++
					break
				}
			}
			right := 0
			for cm := c + 1; cm < len(g[r]); cm++ {
				if g[r][cm].Height < h {
					right++
				} else {
					right++
					break
				}
			}

			if s := above * below * left * right; s > best {
				best = s
			}
		}
	}

	return best
}

func (g Grid) CountVisible() int {
	c := 0
	for _, row := range g {
		for _, t := range row {
			if t.Visible {
				c++
			}
		}
	}
	return c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	g := make(Grid, 0)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]*Tree, 0)
		for i := 0; i < len(line); i++ {
			v, err := strconv.ParseInt(string(line[i]), 10, 64)
			if err != nil {
				panic(err)
			}
			row = append(row, &Tree{Height: v})
		}
		g = append(g, row)
	}

	g.MarkVisible()
	log.Printf("part 1: %v", g.CountVisible())

	log.Printf("part 2: %v", g.ScenicScore())

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
