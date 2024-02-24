package main

import (
	"encoding/json"
	"fmt"
	"go-playground/pkg/xkcd"
	"os"
)

const UrlBase = "https://xkcd.com/"

func main() {
	errors := 0
	comics := make([]*xkcd.Comic, 0)
	var store string

	if len(os.Args) < 2 {
		store = "comics.json"
	} else {
		store = os.Args[1]
	}

	fmt.Println("Storing comics in:", store)
	_, err := os.Stat(store)
	if err == nil {
		fmt.Fprintf(os.Stderr, "Store already exists, skipping index process\n")
		return
	}

	file, err := os.Create(store)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(-1)
	}
	defer file.Close()

	for i := 1; ; i++ {
		url := fmt.Sprintf("%s%d/info.0.json", UrlBase, i)
		comic := xkcd.Fetch(url)

		if comic == nil {
			errors++
			if errors == 3 {
				break
			} else {
				continue
			}
		}

		if errors > 0 {
			errors = 0
		}

		comic.Print(os.Stdout)
		comics = append(comics, comic)
	}

	fmt.Println("Total comics fetched:", len(comics))

	err = json.NewEncoder(file).Encode(comics)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding to JSON: %v\n", err)
		return
	}
}
