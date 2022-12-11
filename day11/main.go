package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type WorryFunc func(a int64) int64

func TimesFunc(val int64) WorryFunc {
	return func(old int64) int64 {
		return old * val
	}
}

func AddFunc(val int64) WorryFunc {
	return func(old int64) int64 {
		return old + val
	}
}

func SquareFunc() WorryFunc {
	return func(old int64) int64 {
		return old * old
	}
}

type Monkey struct {
	Number      int
	Items       []int64
	Operation   WorryFunc
	TestDivisor int64
	TrueTarget  int
	FalseTarget int
	Inspected   int64
}

func (m *Monkey) Recv(item int64) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) HasItems() bool {
	return len(m.Items) > 0
}

func (m *Monkey) Inspect(factor int64) *Route {
	m.Inspected++

	item := m.Items[0]
	m.Items = m.Items[1:]

	worry := m.Operation(item)
	// Part1
	// worry = worry / 3
	// Part2
	worry = worry % factor

	if worry%m.TestDivisor == 0 {
		return &Route{
			Number: m.TrueTarget,
			Item:   worry,
		}
	}

	return &Route{
		Number: m.FalseTarget,
		Item:   worry,
	}
}

type Route struct {
	Number int
	Item   int64
}

var errorDone = errors.New("done")

func NewMonkey(scanner *bufio.Scanner) (*Monkey, error) {
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, " ")
	if len(parts) != 2 && parts[0] != "Monkey" {
		return nil, errorDone
	}

	parts = strings.Split(parts[1], ":")
	num, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, ":")
	parts[1] = strings.TrimSpace(parts[1])
	parts = strings.Split(parts[1], ",")

	items := make([]int64, 0)
	for _, is := range parts {
		is = strings.TrimSpace(is)
		i, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	// Operation
	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, "new = ")
	line = parts[1]

	var wf WorryFunc
	if line == "old * old" {
		wf = SquareFunc()
	} else {
		parts = strings.Split(line, " ")
		operand, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return nil, err
		}
		if parts[1] == "+" {
			wf = AddFunc(operand)
		} else if parts[1] == "*" {
			wf = TimesFunc(operand)
		} else {
			return nil, fmt.Errorf("bad operand")
		}
	}

	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, " ")
	test, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, " ")
	tt, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil {
		return nil, err
	}

	scanner.Scan()
	line = scanner.Text()
	parts = strings.Split(line, " ")
	ft, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Monkey{
		Number:      int(num),
		Items:       items,
		Operation:   wf,
		TestDivisor: test,
		TrueTarget:  int(tt),
		FalseTarget: int(ft),
		Inspected:   0,
	}, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	monkeys := make([]*Monkey, 0)
	// Load all the monkeys
	var m *Monkey
	var err error = nil
	for err == nil {
		m, err = NewMonkey(scanner)
		if err == nil {
			monkeys = append(monkeys, m)
		} else if errors.Is(err, errorDone) {
			break
		} else {
			panic(err)
		}
	}
	for _, m := range monkeys {
		log.Printf("%+v", m)
	}

	// Create a map and find the mod factor
	// product of all divisor tests.
	factor := int64(1)
	monkeyMap := make(map[int]*Monkey)
	for _, m := range monkeys {
		monkeyMap[m.Number] = m
		factor *= m.TestDivisor
	}

	// Run the rounds (part2 values)
	for round := 0; round < 10000; round++ {
		for _, m := range monkeys {
			for m.HasItems() {
				r := m.Inspect(factor)
				monkeyMap[r.Number].Recv(r.Item)
			}
		}
	}

	processed := make([]int64, len(monkeys))
	for i, m := range monkeys {
		fmt.Printf("Monkey %v inspected items %v times\n", m.Number, m.Inspected)
		processed[i] = m.Inspected
	}
	sort.Slice(processed, func(i, j int) bool { return processed[i] >= processed[j] })
	fmt.Printf("answer = %v\n", processed[0]*processed[1])

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
