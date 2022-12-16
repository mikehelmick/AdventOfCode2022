package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
)

type Valve struct {
	name   string
	rate   int
	tunnel []string
}

func LoadValve(s string) *Valve {
	// Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
	s = strings.ReplaceAll(s, "Valve ", "")
	s = strings.ReplaceAll(s, " has flow rate=", ",")
	s = strings.ReplaceAll(s, "; tunnels lead to valves ", ",")
	s = strings.ReplaceAll(s, "; tunnel leads to valve ", ",")
	s = strings.ReplaceAll(s, " ", "")

	parts := strings.Split(s, ",")
	name := parts[0]
	rate := int(straid.AsInt(parts[1]))
	tunnel := make([]string, 0)
	for i := 2; i < len(parts); i++ {
		tunnel = append(tunnel, parts[i])
	}
	return &Valve{
		name:   name,
		rate:   rate,
		tunnel: tunnel,
	}
}

type ValveMap map[string]*Valve

// instead of pre-calculating the all pairs shortest path, just calculate on
// demand and cache the answer. Shared between both parts, so it's not that
// expensive and was faster to write.
var spCache = make(map[string]int)

func shortestPath(pos string, target string, valves ValveMap) int {
	idx := fmt.Sprintf("%v-%v", pos, target)
	if v, ok := spCache[idx]; ok {
		//fmt.Printf("path from %v to %v is %v\n", pos, target, v)
		return v
	}

	visited := make(map[string]bool)
	visited[pos] = true
	wave := []string{pos}
	dist := 0
	for len(wave) > 0 {
		next := make([]string, 0)
		for _, p := range wave {
			if p == target {
				spCache[idx] = dist
				return dist
			}
			for _, cand := range valves[p].tunnel {
				if !visited[cand] {
					next = append(next, cand)
					visited[cand] = true
				}
			}
		}
		wave = next
		dist++
	}
	panic(fmt.Sprintf("%+v", spCache))
}

// Represents a valid solution under the constraints
// The opened slice is an ordered list of the valves opened
type Solution struct {
	opened []string
	flow   int

	cache map[string]bool
}

func NewSolution(path []string, flow int) *Solution {
	op := make([]string, len(path))
	copy(op, path)
	cache := make(map[string]bool)
	for _, v := range path {
		cache[v] = true
	}
	return &Solution{
		opened: op,
		flow:   flow,
		cache:  cache,
	}
}

func (s *Solution) String() string {
	return fmt.Sprintf("%+v -> %v", s.opened, s.flow)
}

// Disjoint returns true if two solutions share no common valves.
func (s *Solution) Disjoint(o *Solution) bool {
	for k := range o.cache {
		if s.cache[k] {
			return false
		}
	}
	return true
}

// dfs does a depth first search attempting to open valves and returns all possible solutions
// based on the starting condition.
func dfs(check map[string]bool, path []string, valves ValveMap, press int, minute int, time int, pos string) []*Solution {
	// out of time, we have a solution.
	if minute > time {
		return []*Solution{NewSolution(path, press)}
	}

	// just moved to pos, open it and account for new flow up to time.
	minute++
	press += ((time - minute) * valves[pos].rate)
	// that was the last possible valve to open, so this is a solution.
	if len(check) == 0 {
		return []*Solution{NewSolution(path, press)}
	}

	// make a defensive copy of the map, otherwise things get weird.
	toOpen := make(map[string]bool)
	for k, v := range check {
		toOpen[k] = v
	}

	sols := make([]*Solution, 0)
	// accumulate all possible solutions when pos is opened and check is left to open.
	for k, _ := range toOpen {
		delete(toOpen, k)
		path = append(path, k)
		newSols := dfs(toOpen, path, valves, press, minute+shortestPath(pos, k, valves), time, k)
		toOpen[k] = true
		path = path[0 : len(path)-1]
		sols = append(sols, newSols...)
	}
	return sols
}

func getSolutions(toOpen map[string]bool, valves ValveMap, time int, pos string) []*Solution {
	check := make(map[string]bool)
	for k, v := range toOpen {
		check[k] = v
	}

	sols := make([]*Solution, 0)
	for k, _ := range toOpen {
		// if we open k first... what's the payoff
		delete(check, k)
		newSols := dfs(check, []string{k}, valves, 0, shortestPath(pos, k, valves), time, k)
		check[k] = true
		sols = append(sols, newSols...)
	}
	return sols
}

func part1(toOpen map[string]bool, valves ValveMap, time int, pos string) int {
	sols := getSolutions(toOpen, valves, time, pos)
	sort.Slice(sols, func(i, j int) bool { return sols[i].flow >= sols[j].flow })
	return sols[0].flow
}

func part2(toOpen map[string]bool, valves ValveMap, time int, pos string) int {
	sols := getSolutions(toOpen, valves, time, pos)
	sort.Slice(sols, func(i, j int) bool { return sols[i].flow >= sols[j].flow })

	// find the two highest (sorted) non overlapping
	max := 0
	for i := 0; i < len(sols); i++ {
		for j := i + 1; j < len(sols); j++ {
			nm := sols[i].flow + sols[j].flow
			// each pair needs to be short circuited. Eventually this terminates fast.
			if nm <= max {
				break
			}
			// avoid the expensive operation if we don't need it.
			if sols[i].Disjoint(sols[j]) {
				max = nm
			}
		}
	}
	return max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	valves := make(ValveMap)
	// Restrict the consideration to non-zero flow vales only.
	toOpen := make(map[string]bool, len(valves))
	for scanner.Scan() {
		line := scanner.Text()
		v := LoadValve(line)
		valves[v.name] = v
		if v.rate > 0 {
			toOpen[v.name] = true
		}
	}
	log.Printf("must open: %+v", toOpen)

	pressure := part1(toOpen, valves, 30, "AA")
	log.Printf("part 1 : %v", pressure)
	log.Printf("part 2 : %v", part2(toOpen, valves, 26, "AA"))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
