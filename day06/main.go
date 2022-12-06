package main

import (
	"bufio"
	"log"
	"os"
)

// Ring is a simple string ring buffer.
type Ring struct {
	buffer []string
	size   int
	pos    int
}

func New(size int) *Ring {
	return &Ring{
		buffer: make([]string, size),
		size:   0,
		pos:    0,
	}
}

func (r *Ring) Append(s string) {
	r.buffer[r.pos] = s
	if r.size < len(r.buffer) {
		r.size++
	}
	r.pos++
	if r.pos == len(r.buffer) {
		r.pos = 0
	}
}

func (r *Ring) PacketStart() bool {
	if r.size < len(r.buffer) {
		return false
	}
	m := make(map[string]struct{})
	for _, v := range r.buffer {
		m[v] = struct{}{}
	}
	return len(m) == len(r.buffer)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	r1 := New(4)
	r2 := New(14)

	r1done, r2done := false, false

	for i := 0; i < len(line); i++ {
		r1.Append(string(line[i]))
		r2.Append(string(line[i]))

		if !r1done && r1.PacketStart() {
			log.Printf("packet: %v\n", i+1)
			r1done = true
		}

		if !r2done && r2.PacketStart() {
			log.Printf("message: %v\n", i+1)
			r2done = true
		}
		if r1done && r2done {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
