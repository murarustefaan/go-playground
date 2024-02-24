package xkcd

import (
	"github.com/h2non/gock"
	"testing"
)

func TestFetch(t *testing.T) {
	defer gock.Off()
	gock.New("https://xkcd.com").
		Get("/1/info.0.json").
		Reply(200).
		JSON(map[string]interface{}{
			"num":        1,
			"day":        "1",
			"month":      "1",
			"year":       "2006",
			"title":      "Barrel - Part 1",
			"transcript": "Transcript",
		})

	tests := []struct {
		name string
		uri  string
		want *Comic
	}{
		{name: "valid", uri: "https://xkcd.com/1/info.0.json", want: &Comic{
			Num:        1,
			Day:        "1",
			Month:      "1",
			Year:       "2006",
			Title:      "Barrel - Part 1",
			Transcript: "Transcript",
		}},
		{name: "invalid", uri: "https://xkcd.com/0/info.0.json", want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fetch(tt.uri); equals(got, tt.want) == false {
				t.Errorf("Fetch() = %v, want %v", got, tt.uri)
			}
		})
	}
}

func equals(a, b *Comic) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Num == b.Num &&
		a.Day == b.Day &&
		a.Month == b.Month &&
		a.Year == b.Year &&
		a.Title == b.Title &&
		a.Transcript == b.Transcript
}
