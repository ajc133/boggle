package main

import (
	"strconv"
	"testing"
)

func isEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestConcatSquares(t *testing.T) {
	testCases := []struct {
		squares  []Square
		expected string
	}{
		{
			squares:  []Square{{"", 0, 0}, {"", 1, 1}},
			expected: "",
		},
		{
			squares:  []Square{{"", 2, 2}, {"z", 3, 3}},
			expected: "z",
		},
		{
			squares:  []Square{{"a", 0, 0}, {"", 1, 1}},
			expected: "a",
		},
		{
			squares:  []Square{{"x", 2, 2}, {"yz", 3, 3}},
			expected: "xyz",
		},
		{
			squares:  []Square{{"xy", 2, 2}, {"z", 3, 3}},
			expected: "xyz",
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := ConcatSquares(tc.squares)
			if got != tc.expected || err != nil {
				t.Fatalf("Failed to ConcatSquares! Got: %q. Want: %q. %v", got, tc.expected, err)
			}
		})
	}
}

func TestPrefixMatcher(t *testing.T) {
	testCases := []struct {
		prefix   string
		words    []string
		expected []string
	}{
		{
			prefix:   "a",
			words:    []string{""},
			expected: []string{},
		},
		{
			prefix:   "a",
			words:    []string{"a"},
			expected: []string{"a"},
		},
		{
			prefix:   "a",
			words:    []string{"a", "ab", "ba", "abc"},
			expected: []string{"a", "ab", "abc"},
		},
		{
			prefix:   "a",
			words:    []string{"b"},
			expected: []string{},
		},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := GetPrefixMatches(tc.prefix, tc.words)
			if !isEqual(got, tc.expected) {
				t.Fatalf("Failed to GetPrefixMatches! Got: %q. Want: %q.", got, tc.expected)
			}
		})
	}

}
