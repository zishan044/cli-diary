package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Entry struct {
	title string
	date  time.Time
	text  string
	tags  []string
}

func main() {
	var newEntry Entry

	fmt.Println("title:")
	fmt.Scanf("%v", &newEntry.title)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("entry:")
	scanned := scanner.Scan()
	if scanned {
		newEntry.text = scanner.Text()
	}

	fmt.Println("tags:")
	scanned = scanner.Scan()
	newEntry.tags = strings.Fields(scanner.Text())

	newEntry.date = time.Now()

	fmt.Println()

	fmt.Println(newEntry.title)
	fmt.Println(newEntry.date)
	fmt.Println(strings.Join(newEntry.tags, " "))
	fmt.Println(newEntry.text)
}
