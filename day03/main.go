package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	priority = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Rucksack struct {
	part1 string
	part2 string
}

func New(line string) *Rucksack {
	l := len(line)
	h := l / 2

	return &Rucksack{
		part1: line[0:h],
		part2: line[h:l],
	}
}

func (r *Rucksack) FullIndex() map[string]struct{} {
	idx := make(map[string]struct{})
	for _, c := range r.part1 {
		idx[string(c)] = struct{}{}
	}
	for _, c := range r.part2 {
		idx[string(c)] = struct{}{}
	}
	return idx
}

func (r *Rucksack) DupeScore() int {
	p1m := make(map[string]struct{})
	for _, c := range r.part1 {
		p1m[string(c)] = struct{}{}
	}

	for _, c := range r.part2 {
		dup := string(c)
		if _, ok := p1m[dup]; ok {
			return strings.Index(priority, dup)
		}
	}
	return 0
}

func Intersection(r1, r2, r3 *Rucksack) int {
	i1 := r1.FullIndex()
	i2 := r2.FullIndex()

	input := fmt.Sprintf("%v%v", r3.part1, r3.part2)
	for _, c := range input {
		char := string(c)

		_, in1 := i1[char]
		_, in2 := i2[char]

		if in1 && in2 {
			return strings.Index(priority, char)
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tot := 0
	part2 := 0
	buf := make([]*Rucksack, 0)
	for scanner.Scan() {
		line := scanner.Text()
		r := New(line)

		log.Printf("%v%v = %v\n", r.part1, r.part2, r.DupeScore())
		tot += r.DupeScore()

		buf = append(buf, r)
		if len(buf) == 3 {
			part2 += Intersection(buf[0], buf[1], buf[2])
			buf = make([]*Rucksack, 0)
		}
	}
	log.Printf("part1 = %v", tot)
	log.Printf("part2 = %v", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
