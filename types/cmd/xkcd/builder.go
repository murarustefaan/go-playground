package main

import (
	"encoding/json"
	"fmt"
	"go-playground/pkg/xkcd"
	"go-playground/pkg/xkcd/builder"
	"os"
)

func CreateStore() {
	errors := 0
	store := xkcd.GetStorePath()

	_, err := os.Stat(store)
	if err == nil {
		fmt.Fprintf(os.Stderr, "Store already exists, skipping download process.\n")
		return
	}

	fmt.Println("Storing comics in:", store)
	file, err := os.Create(store)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(-1)
	}
	defer file.Close()

	comics := make([]*xkcd.Comic, 0)
	for i := 1; ; i++ {
		url := fmt.Sprintf("%s%d/info.0.json", xkcd.UrlBase, i)
		comic := builder.Fetch(url)

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
