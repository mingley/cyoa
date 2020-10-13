package main

import (
	"os"
	"fmt"
	"flag"
	"cyoa"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("using the story in %s. \n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
	
}