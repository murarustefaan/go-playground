package builder

import (
	"encoding/json"
	"fmt"
	"go-playground/pkg/xkcd"
	"net/http"
	"os"
)

func Fetch(route string) *xkcd.Comic {
	fmt.Println("Downloading: " + route)
	response, err := http.Get(route)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading: %v\n", err)
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Error downloading: %v\n", response.Status)
		return nil
	}

	var comic xkcd.Comic
	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling from JSON: %v\n", err)
		return nil
	}

	return &comic
}
