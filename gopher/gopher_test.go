package gopher

import "testing"

func TestPlural(t *testing.T) {
	// Tab containing test cases
	cases := []struct {
		Input  string
		Count  int
		Output string
	}{
		{Input: "tomate", Count: 1, Output: "tomate"},
		{Input: "tomate", Count: 2, Output: "tomates"},
		{Input: "tomate", Count: 3, Output: "tomates"},
		{Input: "tomate", Count: 9999, Output: "tomates"},
	}

	for _, c := range cases {
		result := Plural(c.Input, c.Count)
		if result != c.Output {
			t.Errorf("expected %s, got %s", c.Output, result)
		}
	}
}
