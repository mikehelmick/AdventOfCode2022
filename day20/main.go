package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mikehelmick/AdventOfCode2022/pkg/list"
	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
)

func solve(data list.List[int64], locator list.List[int], multiplier int64, rounds int) int64 {
	if multiplier > 1 {
		for i, d := range data {
			data[i] = d * multiplier
		}
	}

	l := len(data)
	for c := 0; c < rounds; c++ {
		//fmt.Printf("\n----------\nIteration %v\n", c)
		for i := 0; i < l; i++ {
			fromIdx := locator.Find(i)

			value := data[fromIdx]
			var check int64
			check, data = data.RemoveAt(fromIdx)
			if check != value {
				panic("mismatch")
			}
			var locV int
			locV, locator = locator.RemoveAt(fromIdx)

			dest := int((int64(fromIdx) + value) % int64(l-1))
			for dest < 0 {
				dest += (l - 1)
			}
			newPos := int(dest)
			//fmt.Printf("Moving %v from idx: %v to idx: %v\n", value, fromIdx, newPos)
			data = data.Add(newPos, value)
			locator = locator.Add(newPos, locV)

			//fmt.Printf("D: %+v\n", data)
			//fmt.Printf("L: %+v\n\n", locator)
		}
		//fmt.Printf("D: %+v\n", data)
	}

	zeroIdx := data.Find(0)
	log.Printf("%v %v %v %v", zeroIdx, (zeroIdx+1000)%l, (zeroIdx+2000)%l, (zeroIdx+3000)%l)
	sum := data[(zeroIdx+1000)%l] + data[(zeroIdx+2000)%l] + data[(zeroIdx+3000)%l]
	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	data := list.New[int64](0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, straid.AsInt(line))

	}
	l := len(data)
	locator := list.New[int](l)
	for i := 0; i < l; i++ {
		locator = append(locator, i)
	}
	//fmt.Printf("D: %+v\n", data)
	//fmt.Printf("L: %+v\n", locator)
	fmt.Printf("Length %v\n\n", len(data))

	p1d := make([]int64, len(data))
	copy(p1d, data)
	p1l := make([]int, len(locator))
	copy(p1l, locator)
	part1 := solve(p1d, p1l, 1, 1)
	log.Printf("Part 1: %v", part1)

	part2 := solve(data, locator, 811589153, 10)
	log.Printf("Part 2: %v", part2)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
