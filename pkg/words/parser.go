package words

import (
	"golang.org/x/net/html"
	"strings"
)

func ExtractLinesAndWords(text string) (int, int, error) {
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		return 0, 0, err
	}

	words, lines := count(doc)
	return lines, words, nil
}

func count(node *html.Node) (int, int) {
	words := 0
	lines := 0

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		visitedWords, visitedLines := count(c)

		words += visitedWords
		lines += visitedLines
	}

	if node.Type != html.TextNode || IsWhitespace(node.Data) == true {
		return words, lines
	}

	words += len(strings.Fields(node.Data))
	lines += strings.Count(node.Data, "\n") + 1

	return words, lines
}
