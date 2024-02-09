package main

import (
	"bufio"
	"fmt"
	"log"
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

	file, err := os.OpenFile("entries.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("error opening file")
	}

	_, err = file.WriteString(newEntry.title + "\n" + newEntry.date.Format("Jan 2, 2006") + "\n" + strings.Join(newEntry.tags, " ") + "\n" + newEntry.text + "\n")
	if err != nil {
		log.Fatalf("error writing diary")
	}

}
