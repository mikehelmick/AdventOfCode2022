package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type fstype int

const (
	FILE fstype = iota
	DIR
)

type Node struct {
	Name     string
	FSType   fstype
	Size     int64
	Parent   *Node
	Children []*Node

	totalSize int64
}

func NewNode(name string, typ fstype, size int64, parent *Node) *Node {
	return &Node{
		Name:     name,
		FSType:   typ,
		Size:     size,
		Parent:   parent,
		Children: make([]*Node, 0),
	}
}

func (n *Node) AddChild(name string, typ fstype, size int64) *Node {
	child := NewNode(name, typ, size, n)
	n.Children = append(n.Children, child)
	return child
}

func (n *Node) TotalSize() int64 {
	if n.FSType == FILE {
		return n.Size
	}
	if n.totalSize > 0 {
		return n.totalSize
	}

	sum := int64(0)
	for _, c := range n.Children {
		sum += c.TotalSize()
	}
	// memoize because the tree currently can't change.
	n.totalSize = sum
	return sum
}

func (n *Node) Print(depth int) {
	sp := strings.Repeat(" ", depth)
	if n.FSType == FILE {
		log.Printf("%s%s (%v)", sp, n.Name, n.Size)
	} else {
		log.Printf("%s%s DIR (%v)", sp, n.Name, n.TotalSize())
		for _, c := range n.Children {
			c.Print(depth + 1)
		}
	}
}

// Does a DFS, counting
func (n *Node) SumIf(f func(n *Node) bool) int64 {
	sum := int64(0)

	if f(n) {
		sum += n.TotalSize()
	}

	for _, c := range n.Children {
		sum += c.SumIf(f)
	}
	return sum
}

func AllDirs(n *Node) []*Node {
	rtn := make([]*Node, 0)
	if n.FSType == FILE {
		return rtn
	}
	rtn = append(rtn, n)

	for _, c := range n.Children {
		if c.FSType == DIR {
			rtn = append(rtn, c)
			rtn = append(rtn, AllDirs(c)...)
		}
	}
	return rtn
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	root := NewNode("/", DIR, 0, nil)
	cur := root

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "$") {
			// this is a command.
			parts := strings.Split(line, " ")
			if parts[1] == "cd" {
				if parts[2] == "/" {
					cur = root
				} else if parts[2] == ".." {
					if cur.Parent != nil {
						cur = cur.Parent
					}
				} else {
					found := false
					for _, n := range cur.Children {
						if n.Name == parts[2] {
							cur = n
							found = true
							break
						}
					}
					// If we changing into a directory we haven't seen, add it.
					if !found {
						cur = cur.AddChild(parts[2], DIR, 0)
					}
				}
			} else if parts[1] == "ls" {
				// Nothing.
			} else {
				log.Fatalf("unknown command: %v", line)
			}
		} else {
			// we're in a file listing
			parts := strings.Split(line, " ")
			if parts[0] == "dir" {
				cur.AddChild(parts[1], DIR, 0)
				//log.Printf("new dir: %+v", dir)
			} else {
				sz, err := strconv.ParseInt(parts[0], 10, 64)
				if err != nil {
					panic(err)
				}
				cur.AddChild(parts[1], FILE, sz)
				//log.Printf("new file: %+v", file)
			}
		}
	}

	//root.Print(0)
	part1 := root.SumIf(func(n *Node) bool {
		return n.FSType == DIR && n.TotalSize() <= 100000
	})
	log.Printf("part 1 answer: %v", part1)

	spaceNeeded := 30000000 - (int64(70000000) - root.TotalSize())
	log.Printf("Need to free: %v", spaceNeeded)

	allDirs := AllDirs(root)
	sort.Slice(allDirs, func(i, j int) bool { return allDirs[i].TotalSize() < allDirs[j].TotalSize() })
	for _, c := range allDirs {
		if c.TotalSize() >= spaceNeeded {
			log.Printf("part 2 answer: %v", c.TotalSize())
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
