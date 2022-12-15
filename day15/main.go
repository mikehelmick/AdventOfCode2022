package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/twod"
)

func Parse(s string) (*twod.Pos, *twod.Pos) {
	// Sensor at x=2, y=18: closest beacon is at x=-2, y=15
	s = strings.ReplaceAll(s, "Sensor at ", "")
	s = strings.ReplaceAll(s, ": closest beacon is at ", ",")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "x=", "")
	s = strings.ReplaceAll(s, "y=", "")

	parts := strings.Split(s, ",")

	sensor := &twod.Pos{
		Row: int(straid.AsInt(parts[1])),
		Col: int(straid.AsInt(parts[0])),
	}
	beacon := &twod.Pos{
		Row: int(straid.AsInt(parts[3])),
		Col: int(straid.AsInt(parts[2])),
	}

	return sensor, beacon
}

func calcMin(m *twod.Pos, n *twod.Pos) {
	if n.Row < m.Row {
		m.Row = n.Row
	}
	if n.Col < m.Col {
		m.Col = n.Col
	}
}

func calcMax(m *twod.Pos, n *twod.Pos) {
	if n.Row > m.Row {
		m.Row = n.Row
	}
	if n.Col > m.Col {
		m.Col = n.Col
	}
}

type Pair struct {
	sensor *twod.Pos
	beacon *twod.Pos
	dist   int
}

func NewPair(s, b *twod.Pos) *Pair {
	return &Pair{
		sensor: s,
		beacon: b,
		dist:   s.Dist(b),
	}
}

func couldBeBeacon(p *twod.Pos, pairs []*Pair) bool {
	for _, sp := range pairs {
		if p.Dist(sp.sensor) <= sp.dist {
			return false
		}
	}
	return true
}

type Range struct {
	Low  int
	High int
}

func (r *Range) String() string {
	return fmt.Sprintf("[%v - %v]", r.Low, r.High)
}

func NewRange(l, h int) *Range {
	return &Range{Low: l, High: h}
}

// assumes sorted.
func (r *Range) Merge(o *Range) error {
	if o.Low >= r.Low && o.Low <= r.High {
		if o.High > r.High {
			r.High = o.High
		}
		return nil
	}
	return fmt.Errorf("found it")
}

func part2(pairs []*Pair) *twod.Pos {
	// Check each row in the search space.
	for r := 0; r <= 4000000; r++ {
		// figure out which items are covered by each sensor.
		ranges := make([]*Range, 0, len(pairs))
		for _, pair := range pairs {
			dist := pair.dist
			rDist := int(math.Abs(float64(pair.sensor.Row - r)))

			if rDist > dist {
				// that sensor doesn't cover anything in this row.
				continue
			}
			rem := dist - rDist
			// Make a range of all Cols covered by this sensor.
			thisRange := NewRange(pair.sensor.Col-rem, pair.sensor.Col+rem)
			ranges = append(ranges, thisRange)
		}

		// Sort the range.
		sort.Slice(ranges, func(i, j int) bool {
			if ranges[i].Low == ranges[j].Low {
				return ranges[i].High <= ranges[j].High
			}
			return ranges[i].Low <= ranges[j].Low
		})

		// Merge the ranges.
		thisR := ranges[0]
		for i := 1; i < len(ranges); i++ {
			// As soon as one doesn't marge - we found the gap.
			if err := thisR.Merge(ranges[i]); err != nil {
				return &twod.Pos{
					Row: r,
					Col: thisR.High + 1,
				}
			}
		}
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	min := &twod.Pos{Row: 1000000, Col: 1000000}
	max := &twod.Pos{Row: 0, Col: 0}

	pairs := make([]*Pair, 0)

	sensors := make(map[string]bool)
	beacons := make(map[string]bool)

	maxD := 0

	for scanner.Scan() {
		line := scanner.Text()
		s, b := Parse(line)
		calcMin(min, s)
		calcMin(min, b)
		calcMax(max, s)
		calcMax(max, b)
		p := NewPair(s, b)
		log.Printf("%v %v dist: %v", s, b, p.dist)
		pairs = append(pairs, p)

		sensors[s.String()] = true
		beacons[b.String()] = true
		if p.dist > maxD {
			maxD = p.dist
		}
	}
	log.Printf("min: %v max: %v", min, max)

	notCovered := 0
	row := 2000000

	// This can also be used to print the example output :)
	//for row := min.Row; row <= max.Row; row++ {
	for col := min.Col - maxD; col <= max.Col+maxD; col++ {
		p := &twod.Pos{
			Row: row, Col: col,
		}
		if sensors[p.String()] {
			//fmt.Printf("S")
			continue
		} else if beacons[p.String()] {
			//fmt.Printf("B")
			continue
		} else {
			if !couldBeBeacon(p, pairs) {
				//fmt.Printf("#")
				//if row == 10 {
				notCovered++
				//}
			} // else {
			//fmt.Printf(".")
			//}
		}
	}
	fmt.Printf("\n")
	//}
	log.Printf("part 1: %v\n", notCovered)

	p := part2(pairs)
	log.Printf("Candidate: %v", p)
	log.Printf("Part 2: %v", p.Col*4000000+p.Row)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
