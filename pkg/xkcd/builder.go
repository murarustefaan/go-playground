package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Comic struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func Fetch(route string) *Comic {
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

	var comic Comic
	err = json.NewDecoder(response.Body).Decode(&comic)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling from JSON: %v\n", err)
		return nil
	}

	return &comic
}

func (c *Comic) Print(writer io.Writer) {
	fmt.Fprintf(writer, "#%d - %s/%s/%s - %s\n", c.Num, c.Day, c.Month, c.Year, c.Title)
}
