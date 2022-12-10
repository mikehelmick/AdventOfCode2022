package twod

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
