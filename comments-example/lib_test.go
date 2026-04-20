package nameparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNameParser_Parse(t *testing.T) {
	for _, test := range []struct {
		name     string
		input    string
		expected ParsedName
	}{
		{
			name:  "Spanish name",
			input: "Manuel Doncel Martos",
			expected: ParsedName{
				Name:    "Manuel",
				Surname: "Doncel Martos",
			},
		},
		{
			name:  "English name",
			input: "John Doe",
			expected: ParsedName{
				Name:    "John",
				Surname: "Doe",
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			p := NameParser{}
			result := p.Parse(test.input)

			if diff := cmp.Diff(test.expected, result); diff != "" {
				t.Fatalf("unexpected result: %s", diff)
			}
		})
	}
}
