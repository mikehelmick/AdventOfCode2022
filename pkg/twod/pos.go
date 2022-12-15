package twod

import "fmt"

var (
	Dirs = map[string]*Pos{
		"R": NewPos(0, 1),
		"U": NewPos(-1, 0),
		"L": NewPos(0, -1),
		"D": NewPos(1, 0),
	}

	Diags = []*Pos{
		{Row: 1, Col: 1},
		{Row: -1, Col: 1},
		{Row: -1, Col: -1},
		{Row: 1, Col: -1},
	}
)

type Pos struct {
	Row int
	Col int
}

func NewPos(r, c int) *Pos {
	return &Pos{
		Row: r,
		Col: c,
	}
}

func (p *Pos) String() string {
	return fmt.Sprintf("{%v,%v}", p.Row, p.Col)
}

type ValidFunc func(p *Pos) bool

func (p *Pos) Dist(o *Pos) int {
	x := p.Col - o.Col
	if x <= 0 {
		x *= -1
	}
	y := p.Row - o.Row
	if y <= 0 {
		y *= -1
	}
	return x + y
}

func (p *Pos) Equals(o *Pos) bool {
	return p.Row == o.Row && p.Col == o.Col
}

func (p *Pos) Neighbors(f ValidFunc) []*Pos {
	neighbors := make([]*Pos, 0, len(Dirs))
	for _, d := range Dirs {
		n := p.Clone()
		n.Add(d)
		if f(n) {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (p *Pos) Clone() *Pos {
	return &Pos{
		Row: p.Row,
		Col: p.Col,
	}
}

func (p *Pos) Add(o *Pos) {
	p.Row += o.Row
	p.Col += o.Col
}
