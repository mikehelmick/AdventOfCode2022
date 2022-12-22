package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

const (
	WALL = 1
	OPEN = 2
)

var dirs = map[string]*twod.Pos{
	"R": twod.NewPos(0, 1),
	"D": twod.NewPos(1, 0),
	"L": twod.NewPos(0, -1),
	"U": twod.NewPos(-1, 0),
}

var turns = map[string]map[string]string{
	"R": {"R": "D", "L": "U"},
	"D": {"R": "L", "L": "R"},
	"L": {"R": "U", "L": "D"},
	"U": {"R": "R", "L": "L"},
}

var facing = map[string]int{
	"R": 0,
	"D": 1,
	"L": 2,
	"U": 3,
}

type Maze [][]int

func (m Maze) FindStart() *twod.Pos {
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			if m[r][c] == OPEN {
				return twod.NewPos(r, c)
			}
		}
	}
	panic("no starting position")
}

func (m Maze) String() string {
	s := ""
	for _, row := range m {
		for _, cel := range row {
			if cel == 0 {
				s = fmt.Sprintf("%s ", s)
			} else if cel == WALL {
				s = fmt.Sprintf("%s#", s)
			} else if cel == OPEN {
				s = fmt.Sprintf("%s.", s)
			}
		}
		s = fmt.Sprintf("%s\n", s)
	}
	return s
}

func (m Maze) AddRow(s string, w int) Maze {
	row := make([]int, w)
	for i, c := range s {
		char := string(c)
		if char == "#" {
			row[i+1] = WALL
		} else if char == "." {
			row[i+1] = OPEN
		}
	}
	return append(m, row)
}

func (m Maze) FirstOpenInRow(r int, s int, d int) int {
	for i := s; i >= 0 && i < len(m[r]); i += d {
		switch m[r][i] {
		case 0:
			continue
		case WALL:
			return -1
		case OPEN:
			return i
		}
	}
	panic(fmt.Sprintf("cannot wrap row: %v", r))
}

func (m Maze) FirstOpenInCol(c int, s int, d int) int {
	for i := s; i >= 0 && i < len(m); i += d {
		switch m[i][c] {
		case 0:
			continue
		case WALL:
			return -1
		case OPEN:
			return i
		}
	}
	panic(fmt.Sprintf("cannot wrap col: %v", c))
}

func (m Maze) IsOpen(p *twod.Pos) bool {
	return m[p.Row][p.Col] == OPEN
}

func (m Maze) IsWall(p *twod.Pos) bool {
	return m[p.Row][p.Col] == WALL
}

func (m Maze) IsOutOfBounds(p *twod.Pos) bool {
	return m[p.Row][p.Col] == 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := make([]string, 0)
	width := 0
	path := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "L") {
			path = line
		} else {
			lines = append(lines, line)
			if len(line) > width {
				width = len(line)
			}
		}
	}
	width += 2
	maze := make(Maze, 0, len(lines)+2)
	maze = maze.AddRow(" ", width)
	for _, line := range lines {
		maze = maze.AddRow(line, width)
	}
	maze = maze.AddRow(" ", width)

	fmt.Printf("%+v", maze)
	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	parts := strings.Split(path, " ")
	log.Printf("%+v", parts)

	pos := maze.FindStart()
	log.Printf("starting at %+v", pos)

	// pos is final spot?
	part1 := solve(maze, parts, part1wrap)
	log.Printf("part 1: %v", part1)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

type WrapFun func(Maze, string, *twod.Pos) *twod.Pos

var part1wrap = func(maze Maze, dir string, pos *twod.Pos) *twod.Pos {
	switch dir {
	case "R":
		nc := maze.FirstOpenInRow(pos.Row, 0, 1)
		if nc == -1 {
			return pos
		}
		return twod.NewPos(pos.Row, nc)
	case "L":
		nc := maze.FirstOpenInRow(pos.Row, len(maze[0])-1, -1)
		if nc == -1 {
			return pos
		}
		return twod.NewPos(pos.Row, nc)
	case "D":
		nr := maze.FirstOpenInCol(pos.Col, 0, 1)
		if nr == -1 {
			return pos
		}
		return twod.NewPos(nr, pos.Col)
	case "U":
		nr := maze.FirstOpenInCol(pos.Col, len(maze)-1, -1)
		if nr == -1 {
			return pos
		}
		return twod.NewPos(nr, pos.Col)
	}
	panic("you're lost")
}

func solve(maze Maze, parts []string, wrap WrapFun) int {
	pos := maze.FindStart()
	dir := "R"
	for i, p := range parts {
		if i%2 == 0 {
			steps := int(straid.AsInt(p))

			for s := 0; s < steps; s++ {
				next := pos.Clone()
				next.Add(dirs[dir])
				if maze.IsOpen(next) {
					pos = next
				}
				if maze.IsWall(next) {
					break
				}
				if maze.IsOutOfBounds(next) {
					pos = wrap(maze, dir, pos)
				}
			}
		} else {
			dir = turns[dir][p]
		}
	}
	answer := 1000*pos.Row + 4*pos.Col + facing[dir]
	return answer
}
