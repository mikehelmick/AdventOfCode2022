package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
	"golang.org/x/sync/semaphore"
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

func dfs(check map[string]bool, valves ValveMap, press int, minute int, time int, pos string) int {
	toOpen := make(map[string]bool)
	for k, v := range check {
		toOpen[k] = v
	}
	if minute > time {
		return press
	}

	minute++
	//fmt.Printf("%v", strings.Repeat(" ", 10-len(toOpen)))
	added := (time - minute) * valves[pos].rate
	press += added
	//fmt.Printf("open %v at minute %v provides: %v --> %v\n", pos, minute, added, press)

	if len(toOpen) == 0 {
		return press
	}

	max := 0
	for k, _ := range toOpen {
		delete(toOpen, k)
		v := dfs(toOpen, valves, press, minute+shortestPath(pos, k, valves), time, k)
		toOpen[k] = true
		if v > max {
			max = v
		}
	}
	return max
}

func part1(toOpen map[string]bool, valves ValveMap, time int, pos string) int {
	check := make(map[string]bool)
	for k, v := range toOpen {
		check[k] = v
	}

	max := 0
	for k, _ := range toOpen {
		//fmt.Printf("opening %v first\n", k)
		// if we open k first... what's the payoff
		delete(check, k)
		v := dfs(check, valves, 0, shortestPath(pos, k, valves), time, k)
		check[k] = true
		if v > max {
			max = v
		}
	}
	return max
}

func dfs2(toOpen map[string]bool, valves ValveMap, press int, minute int, minuteE int, time int, pos string, epos string) int {
	check := make(map[string]bool)
	for k, v := range toOpen {
		check[k] = v
	}

	if minute > time && minuteE > time {
		return press
	} else if minute > time {
		minuteE++
		added := (time - minuteE) * valves[epos].rate
		press += added
		return dfs(toOpen, valves, press, minuteE, time, epos)
	} else if minuteE > time {
		minute++
		added := (time - minute) * valves[pos].rate
		press += added
		return dfs(toOpen, valves, press, minute, time, pos)
	}

	minute++
	//fmt.Printf("%v", strings.Repeat(" ", 10-len(toOpen)))
	added := (time - minute) * valves[pos].rate
	press += added

	minuteE++
	added = (time - minuteE) * valves[epos].rate
	press += added
	//fmt.Printf("open %v at minute %v provides: %v --> %v\n", pos, minute, added, press)

	if len(toOpen) == 1 {
		for l, _ := range toOpen {
			// who would get there faster
			if minute+shortestPath(pos, l, valves) < minuteE+shortestPath(epos, l, valves) {
				return dfs(map[string]bool{}, valves, press, minute+shortestPath(pos, l, valves), time, l)
			} else {
				return dfs(map[string]bool{}, valves, press, minuteE+shortestPath(epos, l, valves), time, l)
			}
		}
	}

	if len(toOpen) == 0 {
		return press
	}

	max := 0
	for k, _ := range toOpen {
		for e, _ := range toOpen {
			if k == e {
				continue
			}
			delete(check, k)
			delete(check, e)
			v := dfs2(check, valves, press, minute+shortestPath(pos, k, valves), minuteE+shortestPath(epos, e, valves), time, k, e)
			check[k] = true
			check[e] = true
			if v > max {
				max = v
			}
		}
	}
	return max
}

func part2(toOpen map[string]bool, valves ValveMap, time int, pos string) int {
	ctx := context.Background()
	sem := semaphore.NewWeighted(20)
	max := 0
	var mu sync.Mutex
	for k, _ := range toOpen {
		for e, _ := range toOpen {
			k := k
			e := e
			if k == e {
				continue
			}

			check := make(map[string]bool)
			for k, v := range toOpen {
				check[k] = v
			}

			sem.Acquire(ctx, 1)

			go func() {
				fmt.Printf("opening %v and %v first\n", k, e)
				// if we open k first... what's the payoff
				delete(check, k)
				delete(check, e)
				v := dfs2(check, valves, 0, shortestPath(pos, k, valves), shortestPath(pos, e, valves), time, k, e)
				check[k] = true
				check[e] = true
				fmt.Printf("  -> v=%v\n", v)
				mu.Lock()
				defer mu.Unlock()
				if v > max {
					max = v
				}
				sem.Release(1)
			}()
		}
	}

	sem.Acquire(ctx, 20)
	sem.Release(20)

	return max
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	valves := make(ValveMap)
	for scanner.Scan() {
		line := scanner.Text()
		v := LoadValve(line)
		valves[v.name] = v
	}
	log.Printf("%+v", valves)

	toOpen := make(map[string]bool, len(valves))
	for k, v := range valves {
		if v.rate > 0 {
			toOpen[k] = true
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
