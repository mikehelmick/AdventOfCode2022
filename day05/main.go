package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []string
}

func New() *Stack {
	return &Stack{
		data: make([]string, 0),
	}
}

func (s *Stack) Clone() *Stack {
	data := make([]string, len(s.data))
	copy(data, s.data)
	return &Stack{
		data: data,
	}
}

func (s *Stack) Peek() string {
	return string(s.data[len(s.data)-1])
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(str string) {
	s.data = append(s.data, str)
}

func (s *Stack) Pop() string {
	r := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	init := New()
	numStacks := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, " 1") {
			p := strings.Fields(line)
			numStacks = len(p)
			break
		}
		init.Push(line)
	}
	scanner.Scan()
	scanner.Text() // consume empty line

	stacks := make([]Stack, numStacks)
	for !init.Empty() {
		line := init.Pop()
		for i := 1; i <= numStacks; i++ {
			pos := (i-1)*4 + 1
			if s := string(line[pos]); s != " " {
				stacks[i-1].Push(s)
			}
		}
	}

	// For solving p1 and p2 at the same time, deep copy the data.
	p2stacks := make([]*Stack, numStacks)
	for i, s := range stacks {
		p2stacks[i] = s.Clone()
	}

	for scanner.Scan() {
		command := scanner.Text()
		parts := strings.Fields(command)

		amount, _ := strconv.ParseInt(parts[1], 10, 64)
		from, _ := strconv.ParseInt(parts[3], 10, 64)
		to, _ := strconv.ParseInt(parts[5], 10, 64)

		// Part 1
		for i := 0; i < int(amount); i++ {
			val := stacks[from-1].Pop()
			stacks[to-1].Push(val)
		}

		// Part 2
		tmp := New()
		for i := 0; i < int(amount); i++ {
			tmp.Push(p2stacks[from-1].Pop())
		}
		for !tmp.Empty() {
			p2stacks[to-1].Push(tmp.Pop())
		}

		//log.Printf("command: %v\n", command)
		//for _, s := range stacks {
		//	fmt.Printf("%+v\n", s)
		//}
	}

	part1 := ""
	for _, s := range stacks {
		part1 = fmt.Sprintf("%s%s", part1, s.Peek())
	}
	log.Printf("part1: %v\n", part1)

	part2 := ""
	for _, s := range p2stacks {
		part2 = fmt.Sprintf("%s%s", part2, s.Peek())
	}
	log.Printf("part2: %v\n", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
