package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Allows us to compare lists to numbers, etc.
type Unit interface {
	Compare(o Unit) int
}

type List struct {
	Data     []Unit
	Sentinel bool
}

func (l *List) String() string {
	return fmt.Sprintf("%+v", l.Data)
}

func (l *List) Compare(o Unit) int {
	//log.Printf("comparing %v and %v", l, o)
	switch o.(type) {
	case *Number:
		return l.Compare(o.(*Number).ToList())
	case *List:
		// This is the critical piece of the solution.
		right := o.(*List)
		for i, v := range l.Data {
			//log.Printf(" --> %v %v", i, v)

			// right side ran out first
			if i >= len(right.Data) {
				//log.Printf("right ran out")
				return 1
			}
			rv := right.Data[i]
			cmp := v.Compare(rv)
			if cmp == 0 {
				continue
			}
			return cmp
		}
		if len(right.Data) > len(l.Data) {
			return -1
		}
		return 0
	}
	panic("what")
	//return -1
}

type Number struct {
	Value int64
}

func (n *Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *Number) ToList() *List {
	return &List{
		Data: []Unit{
			&Number{
				Value: n.Value,
			},
		},
	}
}

func (n *Number) Compare(o Unit) int {
	switch o.(type) {
	case *Number:
		//log.Printf("comparing %v and %v", n, o)
		return int(n.Value - o.(*Number).Value)
	case *List:
		return n.ToList().Compare(o)
	}
	return -1
}

func Parse(l string) Unit {
	//log.Printf("%v", l)
	l = strings.ReplaceAll(l, "[", "[,")
	l = strings.ReplaceAll(l, "]", ",]")

	parts := strings.Split(l, ",")

	rtn := &List{
		Data: make([]Unit, 0),
	}

	// Ever time we enter list, push on the stack, every time we exit, pop.
	// The sub-lists need to be referenced in the parent first.
	stack := make([]*List, 0)
	stack = append(stack, rtn)

	for i := 1; i < len(parts); i++ {
		c := parts[i]
		if c == "[" {
			newList := &List{
				Data: make([]Unit, 0),
			}
			last := stack[len(stack)-1]
			last.Data = append(last.Data, newList)
			stack = append(stack, newList)
		} else if c == "]" {
			// pop
			stack = stack[0 : len(stack)-1]
		} else if c == "" {
			continue
		} else {
			v, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				panic(err)
			}
			last := stack[len(stack)-1]
			last.Data = append(last.Data, &Number{
				Value: v,
			})
		}
	}
	return rtn
}

type Pair struct {
	Left  Unit
	Right Unit
}

func main() {
	pairs := make([]*Pair, 0)
	packets := make([]Unit, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		left := Parse(scanner.Text())
		scanner.Scan()
		right := Parse(scanner.Text())

		pairs = append(pairs, &Pair{
			Left:  left,
			Right: right,
		})

		packets = append(packets, left)
		packets = append(packets, right)

		scanner.Scan()
	}

	// add the key packets
	packets = append(packets, &List{
		Data:     []Unit{&Number{Value: 2}},
		Sentinel: true,
	})
	packets = append(packets, &List{
		Data:     []Unit{&Number{Value: 6}},
		Sentinel: true,
	})

	log.Printf("# pairs %v", len(pairs))

	s := 0
	for i, p := range pairs {
		//log.Printf("-------------")
		a := p.Left.Compare(p.Right)
		//log.Printf("%v answer %v", i+1, a)
		if a < 0 {
			s += (i + 1)
		}
	}
	log.Printf("part1 1: %v", s)

	// The way the compare function was written made part 2 trivial!
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) < 0
	})
	// Find the key packets (which we tagged) and return the answer.
	part2 := 0
	for i, v := range packets {
		if l, ok := v.(*List); ok {
			if l.Sentinel {
				if part2 == 0 {
					part2 = i + 1
				} else {
					part2 *= (i + 1)
					break
				}
			}
		}
	}
	log.Printf("part1 2: %v", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
