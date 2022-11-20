package goset

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		input []string
		len   int
	}{
		{[]string{}, 0},
		{[]string{""}, 1},
		{[]string{"A"}, 1},
		{[]string{"A", "B"}, 2},
		{[]string{"A", "A"}, 1},
		{[]string{"A", "A", "B", "CDEF"}, 3},
		{[]string{"ALMKL", "asdc", "AScas", "sdcsdc", "A", "B", "CDEF", "asdc", "Asdc"}, 8},
	}
	for i, test := range tests {
		s := NewSet(test.input)
		if l := len(s); l != test.len {
			t.Errorf("FAILED (%v): Length expected %v, got %v", i, test.len, len(s))
		}
	}
}

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		newInput      []string
		containsInput string
		expected      bool
	}{
		{[]string{}, "A", false},
		{[]string{}, "", false},
		{[]string{""}, "A", false},
		{[]string{""}, "", true},
		{[]string{"A"}, "", false},
		{[]string{"A"}, "A", true},
		{[]string{"A", "B"}, "A", true},
		{[]string{"A", "B"}, "B", true},
		{[]string{"A", "B"}, "C", false},
		{[]string{"ABC", "B"}, "A", false},
		{[]string{"ABC", "B"}, "ABC", true},
		{[]string{"ABC", "CDED", "KMLKA", "sadcsd"}, "A", false},
		{[]string{"ABC", "A", "KMLKA", "sadcsd"}, "A", true},
	}

	for i, test := range tests {
		s := NewSet(test.newInput)
		if output := s.Contains(test.containsInput); output != test.expected {
			t.Errorf("FAILED (%v): %v.Contains(%v) expected %v, got %v\n", i, s, test.containsInput, test.expected, output)
		}
		if output := s.Contains(test.containsInput); output != test.expected {
			t.Errorf("FAILED (%v): %v.Contains(%v) expected %v, got %v\n", i, s, test.containsInput, test.expected, output)
		}
		if l := len(s); l != len(test.newInput) {
			t.Errorf("FAILED (%v): Length expected %v, got %v", i, len(test.newInput), len(s))
		}
	}
}

func TestSet_Add(t *testing.T) {
	tests := []struct {
		addInput       string
		containsInput  string
		containsOutput bool
	}{
		{"", "", true},
		{"", "A", false},
		{"A", "", false},
		{"A", "A", true},
		{"A", "B", false},
		{"AC", "A", false},
		{"AVDS", "AVDS", true},
	}
	for i, test := range tests {
		dummySet := NewSet([]string{})
		dummySet.Add(test.addInput)
		if output := dummySet.Contains(test.containsInput); output != test.containsOutput {
			t.Errorf("FAILED (%v): Set.Add(%v), contains %v excpected %v, got %v.", i, test.addInput, test.containsInput, test.containsOutput, output)
		}
	}
}

func TestSet_Remove(t *testing.T) {
	tests := []struct {
		elements []string
		remove   string
	}{
		{[]string{"A", "B", "C"}, "A"},
		{[]string{"A", "B", "C"}, "B"},
		{[]string{"A", "B", "C"}, "C"},
		{[]string{"A", "B", "C"}, "D"},
		{[]string{}, ""},
		{[]string{}, "S"},
		{[]string{""}, ""},
		{[]string{"ABC", "CDE"}, "ABC"},
		{[]string{"ABC", "CDE"}, "LMNO"},
	}

	for i, test := range tests {
		s := NewSet(test.elements)
		s.Remove(test.remove)
		if s.Contains(test.remove) {
			t.Errorf("FAILED (%v): after remove %v, %v contains %v", i, test.remove, s, test.remove)
		}
	}
}

func TestSet_Union(t *testing.T) {
	tests := []struct {
		s1  Set
		s2  Set
		len int
	}{
		{NewSet([]string{}), NewSet([]string{}), 0},
		{NewSet([]string{"A"}), NewSet([]string{}), 1},
		{NewSet([]string{}), NewSet([]string{"A"}), 1},
		{NewSet([]string{""}), NewSet([]string{"A"}), 2},
		{NewSet([]string{""}), NewSet([]string{"A"}), 2},
		{NewSet([]string{""}), NewSet([]string{""}), 1},
		{NewSet([]string{"A"}), NewSet([]string{"A"}), 1},
		{NewSet([]string{"A", "B"}), NewSet([]string{"A"}), 2},
		{NewSet([]string{"A", "B"}), NewSet([]string{"A", "C"}), 3},
		{NewSet([]string{"A", "B", "DCA"}), NewSet([]string{"A", "C", "KSD", "DCA"}), 5},
	}
	for i, test := range tests {
		u := test.s1.Union(test.s2)
		if len(u) != test.len {
			t.Errorf("FAILED (%v): Length expected %v, got %v\n", i, test.len, len(u))
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		s1  Set
		s2  Set
		len int
	}{
		{NewSet([]string{}), NewSet([]string{}), 0},
		{NewSet([]string{""}), NewSet([]string{""}), 1},
		{NewSet([]string{""}), NewSet([]string{"A"}), 0},
		{NewSet([]string{"A"}), NewSet([]string{""}), 0},
		{NewSet([]string{"A"}), NewSet([]string{"A"}), 1},
		{NewSet([]string{"A", "B"}), NewSet([]string{"B"}), 1},
		{NewSet([]string{"A", "B"}), NewSet([]string{"A"}), 1},
		{NewSet([]string{"A", "B"}), NewSet([]string{"B", "A"}), 2},
		{NewSet([]string{"ABC", "BCA", "DEF", "FED"}), NewSet([]string{"ABC", "DEF", "POI", "BCA"}), 3},
		{NewSet([]string{"A", "B", "C"}), NewSet([]string{"B", "C"}), 2},
	}

	for i, test := range tests {
		inter := test.s1.Intersection(test.s2)
		if len(inter) != test.len {
			t.Errorf("FAILED (%v): Length expected %v, got %v\n", i, test.len, len(inter))
		}
	}
}

func TestSet_Difference(t *testing.T) {
	tests := []struct {
		s1    Set
		s1len int
		s2    Set
		s2len int
		len   int
	}{
		{NewSet([]string{}), 0, NewSet([]string{}), 0, 0},
		{NewSet([]string{""}), 1, NewSet([]string{}), 0, 1},
		{NewSet([]string{}), 0, NewSet([]string{""}), 1, 0},
		{NewSet([]string{""}), 1, NewSet([]string{""}), 1, 0},
		{NewSet([]string{"A"}), 1, NewSet([]string{""}), 1, 1},
		{NewSet([]string{"A"}), 1, NewSet([]string{"A"}), 1, 0},
		{NewSet([]string{"A"}), 1, NewSet([]string{"B"}), 1, 1},
		{NewSet([]string{"A", "B"}), 2, NewSet([]string{"B"}), 1, 1},
		{NewSet([]string{"A", "B"}), 2, NewSet([]string{"C"}), 1, 2},
		{NewSet([]string{"A", "B", "C"}), 3, NewSet([]string{"B"}), 1, 2},
		{NewSet([]string{"A", "B", "C"}), 3, NewSet([]string{"D"}), 1, 3},
		{NewSet([]string{"A", "A"}), 1, NewSet([]string{"A"}), 1, 0},
	}

	for i, test := range tests {
		diff := test.s1.Difference(test.s2)
		if len(diff) != test.len {
			t.Errorf("FAILED (%v): Length expected %v, got %v\n", i, test.len, len(diff))
		}
		if len(test.s1) != test.s1len {
			t.Errorf("FAILED (%v): Length os s1 changed from %v to %v\n", i, test.s1len, len(test.s1))
		}
		if len(test.s2) != test.s2len {
			t.Errorf("FAILED (%v): Length os s1 changed from %v to %v\n", i, test.s2len, len(test.s2))
		}
	}
}
