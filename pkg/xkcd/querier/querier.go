package querier

import (
	"go-playground/pkg/xkcd"
	"strings"
)

func FindComicByKeywords(comics []*xkcd.Comic, keywords []string) *xkcd.Comic {
	for idx := range comics {
		matched := true

		for _, keyword := range keywords {
			title := strings.ToLower(comics[idx].Title)
			transcript := strings.ToLower(comics[idx].Transcript)

			if !strings.Contains(title, keyword) && !strings.Contains(transcript, keyword) {
				matched = false
				break
			}
		}

		if !matched {
			continue
		}

		return comics[idx]
	}

	return nil
}
