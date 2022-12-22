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
	FACE = 3
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
			switch cel {
			case 0:
				s = fmt.Sprintf("%s ", s)
			case WALL:
				s = fmt.Sprintf("%s#", s)
			case OPEN:
				s = fmt.Sprintf("%s.", s)
			case 3:
				s = fmt.Sprintf("%s>", s)
			case 4:
				s = fmt.Sprintf("%sV", s)
			case 5:
				s = fmt.Sprintf("%s<", s)
			case 6:
				s = fmt.Sprintf("%s^", s)
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

// These aren't necessary for part 2, but since I had them from part 1...
// might as well reuse.
func (m Maze) FirstOpenInRow(r int, s int, d int) int {
	for i := s; i >= 0 && i < len(m[r]); i += d {
		switch m[r][i] {
		case 0:
			continue
		case WALL:
			return -1
		default:
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
		default:
			return i
		}
	}
	panic(fmt.Sprintf("cannot wrap col: %v", c))
}

func (m Maze) IsOpen(p *twod.Pos) bool {
	return m[p.Row][p.Col] >= OPEN
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

	path = strings.ReplaceAll(path, "L", " L ")
	path = strings.ReplaceAll(path, "R", " R ")
	parts := strings.Split(path, " ")
	//log.Printf("%+v", parts)

	pos := maze.FindStart()
	log.Printf("starting at %+v", pos)

	// pos is final spot?
	part1 := solve(maze, parts, part1wrap)
	log.Printf("part 1: %v", part1)

	part2 := solve(maze, parts, part2wrap)
	log.Printf("part 2: %v", part2)
	//fmt.Printf("%+v", maze)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

type WrapFun func(Maze, string, *twod.Pos) (string, *twod.Pos)

// Is there a more efficient way to write this... probably.
// I did the make a physical cube and map all the transitions approach, and it works.
var part2wrap = func(maze Maze, dir string, pos *twod.Pos) (string, *twod.Pos) {
	switch dir {
	case "R":
		if pos.Row >= 1 && pos.Row <= 50 {
			// moving from right in 2 to left in 5, but upside down; row 1->150; 50->101
			if nc := maze.FirstOpenInRow(151-pos.Row, 101, -1); nc == -1 {
				return dir, pos
			} else {
				return "L", twod.NewPos(151-pos.Row, nc)
			}
		} else if pos.Row >= 51 && pos.Row <= 100 {
			// from right in 3 to up in 2; row 51-> col 101
			if nr := maze.FirstOpenInCol(pos.Row+50, 51, -1); nr == -1 {
				return dir, pos
			} else {
				return "U", twod.NewPos(nr, pos.Row+50)
			}
		} else if pos.Row >= 101 && pos.Row <= 150 {
			// from right in 5 to left in 2, but upside down; row 101->50, 150->1
			if nc := maze.FirstOpenInRow(151-pos.Row, 151, -1); nc == -1 {
				return dir, pos
			} else {
				return "L", twod.NewPos(151-pos.Row, nc)
			}
		} else if pos.Row >= 151 && pos.Row <= 200 {
			// from right in 6 to up in 5; row 151 -> col 51
			if nr := maze.FirstOpenInCol(pos.Row-100, 151, -1); nr == -1 {
				return dir, pos
			} else {
				return "U", twod.NewPos(nr, pos.Row-100)
			}
		}
	case "L":
		if pos.Row >= 1 && pos.Row <= 50 {
			// left in 1 to right in 4, but upside down; row 1->150, 50->101
			if nc := maze.FirstOpenInRow(151-pos.Row, 0, 1); nc == -1 {
				return dir, pos
			} else {
				return "R", twod.NewPos(151-pos.Row, nc)
			}
		} else if pos.Row >= 51 && pos.Row <= 100 {
			// left in 3 to down in 4; row 51->col 1
			if nr := maze.FirstOpenInCol(pos.Row-50, 100, 1); nr == -1 {
				return dir, pos
			} else {
				return "D", twod.NewPos(nr, pos.Row-50)
			}
		} else if pos.Row >= 101 && pos.Row <= 150 {
			// left in 4 to right in 1, but upside down; row 101->50,150->1
			if nc := maze.FirstOpenInRow(151-pos.Row, 0, 1); nc == -1 {
				return dir, pos
			} else {
				return "R", twod.NewPos(151-pos.Row, nc)
			}
		} else if pos.Row >= 151 && pos.Row <= 200 {
			// left to 6 down in 1; row 151 -> col 51
			if nr := maze.FirstOpenInCol(pos.Row-100, 0, 1); nr == -1 {
				return dir, pos
			} else {
				return "D", twod.NewPos(nr, pos.Row-100)
			}
		}
	case "D":
		if pos.Col >= 1 && pos.Col <= 50 {
			// down in 6 to down in side 2, col 1->101
			if nr := maze.FirstOpenInCol(pos.Col+100, 0, 1); nr == -1 {
				return dir, pos
			} else {
				return "D", twod.NewPos(nr, pos.Col+100)
			}
		} else if pos.Col >= 51 && pos.Col <= 100 {
			// down in 5 to left in side 6; col 51 -> row 151
			if nc := maze.FirstOpenInRow(pos.Col+100, 51, -1); nc == -1 {
				return dir, pos
			} else {
				return "L", twod.NewPos(pos.Col+100, nc)
			}
		} else if pos.Col >= 101 {
			// down in 2 to left in side 3; col 101 -> row 51
			if nc := maze.FirstOpenInRow(pos.Col-50, 101, -1); nc == -1 {
				return dir, pos
			} else {
				return "L", twod.NewPos(pos.Col-50, nc)
			}
		}
	case "U":
		if pos.Col >= 1 && pos.Col <= 50 {
			// up in 4 to right in side 3; col 1 -> row 51
			if nc := maze.FirstOpenInRow(pos.Col+50, 50, 1); nc == -1 {
				return dir, pos
			} else {
				return "R", twod.NewPos(pos.Col+50, nc)
			}
		} else if pos.Col >= 51 && pos.Col <= 100 {
			// up in 1 to right in side 6; col 51 -> row 151
			if nc := maze.FirstOpenInRow(pos.Col+100, 0, 1); nc == -1 {
				return dir, pos
			} else {
				return "R", twod.NewPos(pos.Col+100, nc)
			}
		} else if pos.Col >= 101 {
			// up in 2 to up in side 6; col 101 -> col 1
			if nr := maze.FirstOpenInCol(pos.Col-100, 201, -1); nr == -1 {
				return dir, pos
			} else {
				return "U", twod.NewPos(nr, pos.Col-100)
			}
		}
	}
	panic("you're lost")
}

var part1wrap = func(maze Maze, dir string, pos *twod.Pos) (string, *twod.Pos) {
	switch dir {
	case "R":
		nc := maze.FirstOpenInRow(pos.Row, 0, 1)
		if nc == -1 {
			return dir, pos
		}
		return dir, twod.NewPos(pos.Row, nc)
	case "L":
		nc := maze.FirstOpenInRow(pos.Row, len(maze[0])-1, -1)
		if nc == -1 {
			return dir, pos
		}
		return dir, twod.NewPos(pos.Row, nc)
	case "D":
		nr := maze.FirstOpenInCol(pos.Col, 0, 1)
		if nr == -1 {
			return dir, pos
		}
		return dir, twod.NewPos(nr, pos.Col)
	case "U":
		nr := maze.FirstOpenInCol(pos.Col, len(maze)-1, -1)
		if nr == -1 {
			return dir, pos
		}
		return dir, twod.NewPos(nr, pos.Col)
	}
	panic("you're lost")
}

func solve(maze Maze, parts []string, wrap WrapFun) int {
	pos := maze.FindStart()
	dir := "R"
	for i, p := range parts {
		if i%2 == 0 {
			//log.Printf("%v MOVE %v DIR %v", pos, p, dir)
			//fmt.Printf("%+v", maze)
			steps := int(straid.AsInt(p))

			for s := 0; s < steps; s++ {
				maze[pos.Row][pos.Col] = FACE + facing[dir]
				next := pos.Clone()
				next.Add(dirs[dir])
				if maze.IsOpen(next) {
					pos = next
				}
				if maze.IsWall(next) {
					break
				}
				if maze.IsOutOfBounds(next) {
					dir, pos = wrap(maze, dir, pos)
				}

				//fmt.Printf("%+v", maze)
				//time.Sleep(500 * time.Millisecond)
			}
		} else {
			dir = turns[dir][p]
		}
	}
	log.Printf("ended at %v facing %v", pos, dir)
	answer := 1000*pos.Row + 4*pos.Col + facing[dir]
	return answer
}
