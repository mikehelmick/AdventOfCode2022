package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type ElfStash struct {
	calories []int64
}

func New() *ElfStash {
	return &ElfStash{
		calories: make([]int64, 0),
	}
}

func (e *ElfStash) Add(v int64) {
	e.calories = append(e.calories, v)
}

func (e *ElfStash) Sum() (sum int64) {
	sum = 0
	for _, v := range e.calories {
		sum += v
	}
	log.Printf("sum: %+v == %v", e.calories, sum)
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	elves := make([]*ElfStash, 0)

	e := New()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, e)
			e = New()
			continue
		}

		i, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatalf("cann't parse value: %v", err)
		}
		e.Add(i)
	}
	elves = append(elves, e)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Sum() > elves[j].Sum()
	})

	top3 := elves[0].Sum() + elves[1].Sum() + elves[2].Sum()
	log.Printf("Top3: %v", top3)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
