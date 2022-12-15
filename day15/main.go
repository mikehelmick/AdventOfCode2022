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

type RangeList []*Range

func (rl RangeList) Sort() {
	sort.Slice(rl, func(i, j int) bool {
		if rl[i].Low == rl[j].Low {
			return rl[i].High <= rl[j].High
		}
		return rl[i].Low <= rl[j].Low
	})
}

type Range struct {
	Low  int
	High int
}

func (r *Range) Length() int {
	return r.High - r.Low + 2 // inclusive
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

func rowRanges(row int, pairs []*Pair) RangeList {
	ranges := make(RangeList, 0, len(pairs))
	for _, pair := range pairs {
		dist := pair.dist
		rDist := int(math.Abs(float64(pair.sensor.Row - row)))
		if rDist > dist {
			continue
		}

		rem := dist - rDist
		thisRange := NewRange(pair.sensor.Col-rem, pair.sensor.Col+rem)
		ranges = append(ranges, thisRange)
	}
	ranges.Sort()
	return ranges
}

func part1(row int, pairs []*Pair) int {
	ranges := rowRanges(row, pairs)

	// Attempt to merge ranges.
	merged := make(RangeList, 1, len(ranges))
	merged[0] = ranges[0]
	for i := 1; i < len(ranges); i++ {
		err := merged[len(merged)-1].Merge(ranges[i])
		if err != nil {
			merged = append(merged, ranges[i])
		}
	}

	// Count covered in target row.
	answer := 0
	for _, r := range merged {
		answer += r.Length()
	}
	return answer
}

func part2(pairs []*Pair) *twod.Pos {
	// Check each row in the search space.
	for r := 0; r <= 4000000; r++ {
		ranges := rowRanges(r, pairs)

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

	pairs := make([]*Pair, 0)
	for scanner.Scan() {
		line := scanner.Text()
		s, b := Parse(line)
		p := NewPair(s, b)
		// log.Printf("%v %v dist: %v", s, b, p.dist)
		pairs = append(pairs, p)
	}

	part1 := part1(2000000, pairs)
	log.Printf("part 1: %v\n", part1)

	p := part2(pairs)
	log.Printf("Candidate: %v", p)
	log.Printf("Part 2: %v", p.Col*4000000+p.Row)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
