package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid search keywords provided\n")
		os.Exit(-1)
	}

	keywords := os.Args[1:]
	fmt.Println("Keywords: ", keywords)

	CreateStore()
	comic := RunQuery(keywords)

	comic.Print(os.Stdout)
	fmt.Fprintf(os.Stdout, "%s\n", comic.Transcript)
}
