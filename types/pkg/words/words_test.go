package words

import "testing"

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		name string
		line string
		want bool
	}{
		{"empty", "", true},
		{"spaces", "   ", true},
		{"tabs", "\t\t\t", true},
		{"mixed", "  \t  ", true},
		{"text", "hello", false},
		{"mixed", "  hello  ", false},
		{"mixed", ".  ", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWhitespace(tt.line); got != tt.want {
				t.Errorf("IsWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}
