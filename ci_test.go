package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HI I am LAte toDAy  ",
			expected: []string{"hi", "i", "am", "late", "today"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Fail, lengths of slices dont match: actual:%v expected:%v", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("Fail, words dont match: actual:%v expected:%v", word, expectedWord)
			}
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
