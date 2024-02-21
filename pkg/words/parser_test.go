package words

import (
	"testing"
)

func TestExtractLinesAndWords(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		words int
		lines int
	}{
		{"empty", "", 0, 0},
		{"text one line one word", "hello", 1, 1},
		{"text one line multiple words", "hello world", 2, 1},
		{"text two lines", "hello\nworld", 2, 2},
		{"html", "<html><body>hello\nworld</body></html>", 2, 2},
		{"html multiple tags", "<html><body>world<div>test</div></body></html>", 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, gotWords, err := ExtractLinesAndWords(tt.text)
			if err != nil {
				t.Errorf("ExtractLinesAndWords() error = %v", err)
				return
			}
			if gotWords != tt.words {
				t.Errorf("ExtractLinesAndWords() gotWords = %v, want %v", gotWords, tt.words)
			}
			if gotLines != tt.lines {
				t.Errorf("ExtractLinesAndWords() gotLines = %v, want %v", gotLines, tt.lines)
			}
		})
	}
}
