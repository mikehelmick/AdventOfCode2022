package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	X     int64
	Cycle int
}

func New() *CPU {
	return &CPU{
		X:     1,
		Cycle: 1,
	}
}

func (c *CPU) Tick() {
	c.Cycle++
}

func (c *CPU) Add(val int64) {
	c.X += val
}

func (c *CPU) CheckSignalX() int64 {
	if c.Cycle == 20 || (c.Cycle-20)%40 == 0 {
		return int64(c.Cycle) * c.X
	}
	return 0
}

func (c *CPU) Print() {
	loc := int64(c.Cycle)
	for loc > 40 {
		loc = loc - 40
	}
	loc--

	x := c.X
	if x >= loc-1 && x <= loc+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(" ") // easier to read w/ space.
	}
	if c.Cycle > 1 && (c.Cycle)%40 == 0 {
		fmt.Printf("\n")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var part1 int64
	c := New()

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		if parts[0] == "noop" {
			c.Print()
			c.Tick()
			part1 += c.CheckSignalX()

		} else {
			// must be an add which is 2 ticks, and then add.
			c.Print()
			c.Tick()

			part1 += c.CheckSignalX()
			val, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				panic(err)
			}
			c.Print()
			c.Tick()
			c.Add(val)

			part1 += c.CheckSignalX()
		}
	}

	log.Printf("part 1: %v", part1)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
