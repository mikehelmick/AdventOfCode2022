package main

import (
	"bufio"
	"log"
	"os"
)

func Convert(s string) int64 {
	sum := int64(0)
	place := int64(1)
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		switch r {
		case '0':
		case '1':
			sum += place
		case '2':
			sum += (2 * place)
		case '-':
			sum -= place
		case '=':
			sum -= (2 * place)
		}
		//log.Printf("place: %v r: %v sum: %v", place, string(r), sum)
		place *= 5
	}
	return sum
}

func Reverse(v int64) string {
	if v == 0 {
		return ""
	}

	mod := int(v % 5)
	switch mod {
	case 0:
		return Reverse(v/5) + "0"
	case 1:
		return Reverse(v/5) + "1"
	case 2:
		return Reverse(v/5) + "2"
	case 3:
		return Reverse((v+2)/5) + "="
	case 4:
		return Reverse((v+1)/5) + "-"
	}
	panic("invalid")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		val := Convert(line)
		sum += val
		log.Printf("%v %v", line, val)
	}

	snafu := Reverse(sum)
	log.Printf("Part 1: %v", snafu)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
