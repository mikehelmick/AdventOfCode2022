package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Begin int64
	End   int64
}

func New(in string) (*Range, error) {
	parts := strings.Split(in, "-")

	begin, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse part 0: %w", err)
	}
	end, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cannot parse part 1: %w", err)
	}
	return &Range{
		Begin: begin,
		End:   end,
	}, nil
}

func (r *Range) Contains(o *Range) bool {
	return r.Begin <= o.Begin && r.End >= o.End
}

func (r *Range) Overlaps(o *Range) bool {
	// Order by beginning section, makes comparison easier.
	r1 := r
	r2 := o
	if r2.Begin < r1.Begin {
		r1, r2 = r2, r1
	}
	return r2.Begin <= r1.End
}

func (r *Range) String() string {
	return fmt.Sprintf("%v-%v", r.Begin, r.End)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	contains := 0
	overlaps := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		r1, err := New(parts[0])
		if err != nil {
			panic(err)
		}
		r2, err := New(parts[1])
		if err != nil {
			panic(err)
		}

		if r1.Contains(r2) || r2.Contains(r1) {
			contains++
		}
		if r1.Overlaps(r2) {
			overlaps++
		}
	}

	log.Printf("Fully contains: %v\n", contains)
	log.Printf("Overlaps: %v\n", overlaps)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
