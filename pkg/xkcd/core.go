package xkcd

import (
	"fmt"
	"io"
)

const UrlBase = "https://xkcd.com/"
const DefaultStorePath = "comics.json"

type Comic struct {
	Num        int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func (c *Comic) Print(writer io.Writer) {
	fmt.Fprintf(writer, "#%d - %s/%s/%s - %s\n", c.Num, c.Day, c.Month, c.Year, c.Title)
}

func GetStorePath() string {
	return DefaultStorePath
}
