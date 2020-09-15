package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story %s.\n", *filename)
}
