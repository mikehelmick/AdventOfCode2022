package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
)

type Operator int

const (
	NONE Operator = iota
	ADD
	SUB
	MUL
	DIV
)

var opMap = map[string]Operator{
	"+": ADD,
	"-": SUB,
	"*": MUL,
	"/": DIV,
}

type Element struct {
	Name  string
	Value int64

	Opp   Operator
	Left  *Element
	Right *Element

	LeftName  string
	RightName string
}

func (e *Element) Check() (string, int64) {
	d := e.Left.Calculate() - e.Right.Calculate()
	return fmt.Sprintf("%v == %v (%v)", e.Left.Calculate(), e.Right.Calculate(), d), d
}

func (e *Element) Calculate() int64 {
	switch e.Opp {
	case NONE:
		return e.Value
	case ADD:
		return e.Left.Calculate() + e.Right.Calculate()
	case SUB:
		return e.Left.Calculate() - e.Right.Calculate()
	case MUL:
		return e.Left.Calculate() * e.Right.Calculate()
	case DIV:
		return e.Left.Calculate() / e.Right.Calculate()
	}
	panic("no op")
}

func Load(s string) *Element {
	parts := strings.Split(s, ": ")

	name := parts[0]
	if strings.Contains(parts[1], " ") {
		parts := strings.Split(parts[1], " ")
		return &Element{
			Name:      name,
			Opp:       opMap[parts[1]],
			LeftName:  parts[0],
			RightName: parts[2],
		}
	}
	return &Element{
		Name:  name,
		Value: straid.AsInt(parts[1]),
		Opp:   NONE,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	elem := make([]*Element, 0)
	eMap := make(map[string]*Element)
	for scanner.Scan() {
		line := scanner.Text()
		e := Load(line)
		elem = append(elem, e)
		eMap[e.Name] = e
	}
	// Go through and link the tree.
	for _, v := range elem {
		if v.LeftName != "" {
			v.Left = eMap[v.LeftName]
			v.Right = eMap[v.RightName]
		}
	}

	part1 := eMap["root"].Calculate()
	log.Printf("part 1: %v", part1)

	// This runs almost as fast at 0, but some test runs helped narrow it down.
	low := int64(1000000000000)
	high := int64(8000000000000)
	// Binary search for the humn value that causes check to be zero.
	for low <= high {
		//log.Printf("-- %v -- %v --", low, high)
		median := (low + high) / 2
		eMap["humn"].Value = median

		res, d := eMap["root"].Check()
		//log.Printf("check: %v diff: %v", res, d)
		if d == 0 {
			log.Printf("check: %v", res)
			log.Printf("part 2 : %v", median)
			break
		}
		if d > 0 {
			low = median + 1
		} else {
			high = median - 1
		}

	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
