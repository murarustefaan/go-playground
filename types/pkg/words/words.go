package words

import (
	"regexp"
	"strings"
)

var ignored = [...]string{
	"the", "and", "of", "to", "a", "i", "it", "in", "or", "is", "as", "so", "on", "but", "be", "at", "by", "an", "if", "no", "we", "us", "am", "do", "up", "my", "me", "he", "hi", "go", "is", "it", "in", "on", "or", "to", "so", "as", "at", "be", "by", "if", "of", "an", "am", "we", "us", "do", "up", "he", "hi", "go", "my", "me",
}

type KV struct {
	Key   string
	Value int
}

func NormalizeWord(word string) (string, bool) {
	normalized := strings.ToLower(word)
	normalized = regexp.MustCompile(`[[:punct:]]`).ReplaceAllString(normalized, "")

	if len(normalized) == 0 {
		return "", false
	}

	for _, ignore := range ignored {
		if ignore == normalized {
			return "", false
		}
	}

	return normalized, true
}

func IsWhitespace(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}
