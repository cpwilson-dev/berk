package utils

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input string
		want  []string
	}{
		{
			input: "hello world",
			want: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO WoRld",
			want: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {

		got := CleanInput(cs.input)

		if len(got) != len(cs.want) {
			t.Errorf("The lengths are not equal. Got: %v, want: %v", len(got), len(cs.want))
			continue
		}

		for i := range got {
			gotWord := got[i]
			wantedWord := cs.want[i]

			if gotWord != wantedWord {
				t.Errorf("%v does not equal %v", gotWord, wantedWord)
			}
		}

	}
}
