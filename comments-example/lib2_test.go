package nameparser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNameParser_Parse_2(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input string
		want  ParsedName
	}{
		"Spanish name": {
			input: "Manuel Doncel Martos",
			want: ParsedName{
				Name:    "Manuel",
				Surname: "Doncel Martos",
			},
		},
		"English name": {
			input: "John Doe",
			want: ParsedName{
				Name:    "John",
				Surname: "Doe",
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p := Parser{}
			got := p.Parse(test.input)

			if diff := cmp.Diff(got, test.want); diff != "" {
				t.Errorf("Parse(%q) = %v, diff %v", test.input, got, diff)
			}
		})
	}
}
