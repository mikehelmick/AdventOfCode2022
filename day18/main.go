package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/mathaid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/threed"
)

type Cube struct {
	pos      *threed.Pos
	adjacent []*Cube
}

func NewCube(x, y, z int) *Cube {
	return &Cube{
		pos:      threed.NewPos(x, y, z),
		adjacent: make([]*Cube, 6),
	}
}

func (p *Cube) Pos() *threed.Pos {
	return p.pos.Clone()
}

func (p *Cube) Visible() int {
	c := 0
	for _, b := range p.adjacent {
		if b == nil {
			c++
		}
	}
	return c
}

func (p *Cube) String() string {
	s := make([]string, 6)
	for i, a := range p.adjacent {
		if a == nil {
			s[i] = "nil"
		} else {
			s[i] = a.pos.String()
		}
	}
	return fmt.Sprintf("%s->%+v ", p.pos, s)
}

func (p *Cube) Adjacent(o *Cube) {
	pp := p.pos
	op := o.pos
	if pp.Y == op.Y && pp.Z == op.Z {
		if pp.X-1 == op.X {
			p.adjacent[0] = o
			o.adjacent[1] = p
		} else if pp.X+1 == op.X {
			p.adjacent[1] = o
			o.adjacent[0] = p
		}
	}
	if pp.Y == op.Y && pp.X == op.X {
		if pp.Z-1 == op.Z {
			p.adjacent[4] = o
			o.adjacent[5] = p
		} else if pp.Z+1 == op.Z {
			p.adjacent[5] = o
			o.adjacent[4] = p
		}
	}
	if pp.Z == op.Z && pp.X == op.X {
		if pp.Y-1 == op.Y {
			p.adjacent[2] = o
			o.adjacent[3] = p
		} else if pp.Y+1 == op.Y {
			p.adjacent[3] = o
			o.adjacent[2] = p
		}
	}
}

func (p *Cube) UpdateBounds(bounds []int) {
	bounds[0] = mathaid.Min(p.pos.X, bounds[0])
	bounds[1] = mathaid.Max(p.pos.X, bounds[1])
	bounds[2] = mathaid.Min(p.pos.Y, bounds[2])
	bounds[3] = mathaid.Max(p.pos.Y, bounds[3])
	bounds[4] = mathaid.Min(p.pos.Z, bounds[4])
	bounds[5] = mathaid.Max(p.pos.Z, bounds[5])
}

func enlarge(bounds []int) {
	for i, v := range bounds {
		if i%2 == 0 {
			bounds[i] = v - 1
		} else {
			bounds[i] = v + 1
		}
	}
}

// does a BFS from opposite corners counting each time we run into
// a cube surface.
func cubeBFS(bounds []int, m map[string]bool) int {
	validF := func(p *threed.Pos) bool {
		return p.X >= bounds[0] && p.X <= bounds[1] &&
			p.Y >= bounds[2] && p.Y <= bounds[3] &&
			p.Z >= bounds[4] && p.Z <= bounds[5]
	}

	wave := []*threed.Pos{
		threed.NewPos(bounds[0], bounds[2], bounds[4]),
		threed.NewPos(bounds[1], bounds[3], bounds[5]),
	}
	visited := make(map[string]bool)
	for _, w := range wave {
		visited[w.String()] = true
	}

	count := 0
	for len(wave) > 0 {
		next := make([]*threed.Pos, 0)
		for _, p := range wave {
			neighbors := p.Neighbors(validF)
			for _, cand := range neighbors {
				cs := cand.String()
				if m[cs] {
					// hit a cube
					count++
					visited[cs] = true
					continue
				}
				if !visited[cs] {
					visited[cs] = true
					next = append(next, cand)
				}
			}
		}
		wave = next
	}
	return count
}

func Load(s string) *Cube {
	parts := strings.Split(s, ",")
	return NewCube(
		int(straid.AsInt(parts[0])),
		int(straid.AsInt(parts[1])),
		int(straid.AsInt(parts[2])),
	)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cubes := make([]*Cube, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cubes = append(cubes, Load(line))
	}

	for i := 0; i < len(cubes); i++ {
		for j := i + 1; j < len(cubes); j++ {
			cubes[i].Adjacent(cubes[j])
		}
	}

	cubeMap := make(map[string]bool)

	bounds := []int{100, 0, 100, 0, 100, 0}
	sides := 0
	for _, c := range cubes {
		sides += c.Visible()
		c.UpdateBounds(bounds)
		fmt.Printf("%v\n", c)
		cubeMap[c.Pos().String()] = true
	}
	log.Printf("part 1: %v", sides)

	// adjust bounds out by 1 to make sure we can hit all cubes.
	//log.Printf("bounds %+v", bounds)
	enlarge(bounds)
	//log.Printf("search space %+v", bounds)

	part2 := cubeBFS(bounds, cubeMap)
	log.Printf("part 2: %v", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
