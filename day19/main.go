package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/mikehelmick/AdventOfCode2022/pkg/mathaid"
	"github.com/mikehelmick/AdventOfCode2022/pkg/straid"
)

type Resource int

const (
	NONE Resource = iota
	ORE
	CLAY
	OBSIDIAN
	GEODE
)

var Types = []Resource{ORE, CLAY, OBSIDIAN, GEODE}

var orders = [][]int{
	{1, 2, 3, 4},
	{2, 1, 3, 4},
	{3, 1, 2, 4},
	{1, 3, 2, 4},
	{2, 3, 1, 4},
	{3, 2, 1, 4},
	{3, 2, 4, 1},
	{2, 3, 4, 1},
	{4, 3, 2, 1},
	{3, 4, 2, 1},
	{2, 4, 3, 1},
	{4, 2, 3, 1},
	{4, 1, 3, 2},
	{1, 4, 3, 2},
	{3, 4, 1, 2},
	{4, 3, 1, 2},
	{1, 3, 4, 2},
	{3, 1, 4, 2},
	{2, 1, 4, 3},
	{1, 2, 4, 3},
	{4, 2, 1, 3},
	{2, 4, 1, 3},
	{1, 4, 2, 3},
	{4, 1, 2, 3},
}

type Blueprint struct {
	Number int
	Costs  map[Resource]map[Resource]int
}

func (bp *Blueprint) String() string {
	return fmt.Sprintf("Blueprint: %v Ore: %+v Clay: %+v Obsidian: %+v Geode: %+v",
		bp.Number, bp.Costs[ORE], bp.Costs[CLAY], bp.Costs[OBSIDIAN], bp.Costs[GEODE])
}

func NewBlueprint() *Blueprint {
	bp := &Blueprint{
		Number: 0,
		Costs:  make(map[Resource]map[Resource]int),
	}
	for _, t := range Types {
		bp.Costs[t] = make(map[Resource]int)
	}
	return bp
}

func Load(s string) *Blueprint {
	bp := NewBlueprint()
	// Blueprint 1: Each ore robot costs 3 ore.    // 0,1
	// Each clay robot costs 4 ore.                // 2
	// Each obsidian robot costs 2 ore and 20 clay.// 3, 4
	// Each geode robot costs 4 ore and 7 obsidian.// 5, 6
	s = strings.ReplaceAll(s, "Blueprint ", "")
	s = strings.Replace(s, ": Each ore robot costs ", ",", 1)
	s = strings.Replace(s, " ore. Each clay robot costs ", ",", 1)
	s = strings.Replace(s, " ore. Each obsidian robot costs ", ",", 1)
	s = strings.Replace(s, " ore and ", ",", 1)
	s = strings.Replace(s, " clay. Each geode robot costs ", ",", 1)
	s = strings.Replace(s, " ore and ", ",", 1)
	s = strings.Replace(s, " obsidian.", "", 1)

	p := strings.Split(s, ",")
	bp.Number = int(straid.AsInt(p[0]))
	bp.Costs[ORE][ORE] = int(straid.AsInt(p[1]))
	bp.Costs[CLAY][ORE] = int(straid.AsInt(p[2]))
	bp.Costs[OBSIDIAN][ORE] = int(straid.AsInt(p[3]))
	bp.Costs[OBSIDIAN][CLAY] = int(straid.AsInt(p[4]))
	bp.Costs[GEODE][ORE] = int(straid.AsInt(p[5]))
	bp.Costs[GEODE][OBSIDIAN] = int(straid.AsInt(p[6]))

	return bp
}

type State struct {
	Minute         int
	OreRobots      int
	ClayRobots     int
	ObsidianRobots int
	GeodeRobots    int
	Ore            int
	Clay           int
	Obsidian       int
	Geode          int

	newOre      int
	newClay     int
	newObsidian int
	newGeode    int
}

func (s *State) String() string {
	return fmt.Sprintf("{M:%v OR:%v CR:%v OBR:%b GR:%b O:%v C:%v OB:%v G:%v}",
		s.Minute, s.OreRobots, s.ClayRobots, s.ObsidianRobots, s.GeodeRobots, s.Ore, s.Clay, s.Obsidian, s.Geode)
}

func NewState() *State {
	return &State{
		Minute:    0,
		OreRobots: 1,
	}
}

func (s *State) Clone() *State {
	return &State{
		Minute:         s.Minute,
		OreRobots:      s.OreRobots,
		ClayRobots:     s.ClayRobots,
		ObsidianRobots: s.ObsidianRobots,
		GeodeRobots:    s.GeodeRobots,
		Ore:            s.Ore,
		Clay:           s.Clay,
		Obsidian:       s.Obsidian,
		Geode:          s.Geode,
		newOre:         s.newOre,
		newClay:        s.newClay,
		newObsidian:    s.newObsidian,
		newGeode:       s.newGeode,
	}
}

func (s *State) Tick() {
	s.Minute++
	// save aside, we can't spend it this minute.
	s.newOre = s.OreRobots
	s.newClay = s.ClayRobots
	s.newObsidian = s.ObsidianRobots
	s.newGeode = s.GeodeRobots
}

func (s *State) Save() {
	s.Ore += s.newOre
	s.Clay += s.newClay
	s.Obsidian += s.newObsidian
	s.Geode += s.newGeode
}

func (s *State) Purchase(r Resource, bp *Blueprint) bool {
	switch r {
	case ORE:
		if cost := bp.Costs[ORE][ORE]; s.Ore >= cost {
			s.OreRobots++
			s.Ore -= cost
			return true
		}
	case CLAY:
		if cost := bp.Costs[CLAY][ORE]; s.Ore >= cost {
			s.ClayRobots++
			s.Ore -= cost
			return true
		}
	case OBSIDIAN:
		oreCost := bp.Costs[OBSIDIAN][ORE]
		clayCost := bp.Costs[OBSIDIAN][CLAY]
		if s.Ore >= oreCost && s.Clay >= clayCost {
			s.ObsidianRobots++
			s.Ore -= oreCost
			s.Clay -= clayCost
			return true
		}
	case GEODE:
		oreCost := bp.Costs[GEODE][ORE]
		obsidianCost := bp.Costs[GEODE][OBSIDIAN]
		if s.Ore >= oreCost && s.Obsidian >= obsidianCost {
			s.GeodeRobots++
			s.Ore -= oreCost
			s.Obsidian -= obsidianCost
			return true
		}
	}
	return false
}

func (bp *Blueprint) Enumerate(s *State) []*State {
	// enumerate all possible decisions we could make
	next := make([]*State, 0, 20)
	// Do nothing as an option
	next = append(next, s.Clone())

	for _, order := range orders {
		ns := s.Clone()
		boughtAny := false
		for _, resource := range order {
			if ns.Purchase(Resource(resource), bp) {
				boughtAny = true
				break // This is why my solution didn't work. I let you buy more than 1 robot on a minute...
			}
		}
		if boughtAny {
			next = append(next, ns)
		}
	}
	return next
}

func search(bp *Blueprint, minutes int) int {
	states := []*State{NewState()}

	for i := 1; i <= minutes; i++ {
		log.Printf("Blueprint %v Minute %v States: %v", bp.Number, i, len(states))
		nextStates := make(map[string]*State)
		for _, state := range states {
			state.Tick()
			for _, ns := range bp.Enumerate(state) {
				ns.Save()
				nextStates[ns.String()] = ns
			}
		}
		//log.Printf("BEFORE: %+v", states)
		states = make([]*State, 0, len(nextStates))
		for _, v := range nextStates {
			states = append(states, v)
		}

		if i >= 5 {
			// start culling states that are falling behind on geode robot production
			sort.Slice(states, func(i, j int) bool {
				return states[i].GeodeRobots >= states[j].GeodeRobots
			})
			ret := 0
			needsAtLeast := mathaid.Max[int](0, states[0].GeodeRobots-1)
			for ; ret < len(states); ret++ {
				if states[ret].GeodeRobots < needsAtLeast {
					break
				}
			}
			if ret > 0 {
				states = states[0:ret]
			}
		}

		//log.Printf("%+v", states)
	}

	sort.Slice(states, func(i, j int) bool {
		return states[i].Geode >= states[j].Geode
	})
	return states[0].Geode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	data := make([]*Blueprint, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, Load(line))
	}
	log.Printf("%+v", data)

	total := 0
	for _, bp := range data {
		a := search(bp, 24)
		log.Printf("Blueprint %v has %v geods", bp.Number, a)
		total += (bp.Number * a)
	}
	log.Printf("part 1: %v", total)

	total = 1
	data = data[0:3]
	// not super efficient - since we don't cache the states from part 1
	// but it gets the job done.
	for _, bp := range data {
		a := search(bp, 32)
		log.Printf("Blueprint %v has %v geods", bp.Number, a)
		total *= a
	}
	log.Printf("part 1: %v", total)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
