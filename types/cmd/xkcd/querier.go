package main

import (
	"encoding/json"
	"fmt"
	"go-playground/pkg/xkcd"
	"go-playground/pkg/xkcd/querier"
	"os"
	"strings"
)

func RunQuery(keywords []string) *xkcd.Comic {
	store := xkcd.GetStorePath()

	for idx := range keywords {
		keywords[idx] = strings.ToLower(keywords[idx])
	}

	file, err := os.Open(store)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		return nil
	}

	defer file.Close()

	var comics []*xkcd.Comic
	err = json.NewDecoder(file).Decode(&comics)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding from JSON: %v\n", err)
		return nil
	}

	comic := querier.FindComicByKeywords(comics, keywords)
	if comic == nil {
		fmt.Fprintf(os.Stderr, "No comics found\n")
		return nil
	}

	return comic
}
