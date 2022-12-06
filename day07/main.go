package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("%v\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
