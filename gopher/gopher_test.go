package gopher

import "testing"

func GetTestCases() []struct {
	Input  string
	Count  int
	Output string
} {
	return []struct {
		Input  string
		Count  int
		Output string
	}{
		{Input: "tomate", Count: 1, Output: "tomate"},
		{Input: "tomate", Count: 2, Output: "tomates"},
		{Input: "tomate", Count: 3, Output: "tomates"},
		{Input: "tomate", Count: 9999, Output: "tomates"},
	}
}

func TestPlural(t *testing.T) {
	// Tab containing test cases
	cases := GetTestCases()

	for _, c := range cases {
		result := Plural(c.Input, c.Count)
		if result != c.Output {
			t.Errorf("expected %s, got %s", c.Output, result)
		}
	}
}

func TestPluralWithStringsBuilder(t *testing.T) {
	cases := GetTestCases()

	for _, c := range cases {
		result := PluralWithStringsBuilder(c.Input, c.Count)
		if result != c.Output {
			t.Errorf("expected %s, got %s", c.Output, result)
		}
	}
}
