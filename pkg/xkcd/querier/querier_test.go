package querier

import (
	"go-playground/pkg/xkcd"
	"testing"
)

func TestFindComicByKeywords(t *testing.T) {
	comics := []*xkcd.Comic{
		{Title: "Test 1", Transcript: "Transcript 1"},
		{Title: "Test 2", Transcript: "Transcript 2"},
	}

	tests := []struct {
		name     string
		keywords []string
		want     *xkcd.Comic
	}{
		{name: "valid - 1", keywords: []string{"test"}, want: comics[0]},
		{name: "valid - 2", keywords: []string{"test", "2"}, want: comics[1]},
		{name: "invalid", keywords: []string{"invalid"}, want: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindComicByKeywords(comics, tt.keywords); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
