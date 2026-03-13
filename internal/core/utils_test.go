package core_test

import (
	"docpile/internal/core"
	"testing"
)

func TestCoalesce(t *testing.T) {
	t.Run("test strings", func(t *testing.T) {
		tests := []struct {
			name     string
			values   []string
			expected string
		}{
			{name: "all set", values: []string{"1", "2", "3"}, expected: "1"},
			{name: "last two set", values: []string{"", "2", "3"}, expected: "2"},
			{name: "last one set", values: []string{"", "", "3"}, expected: "3"},
			{name: "all empty", values: []string{"", "", ""}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := core.Coalesce(test.values...)
				if got != test.expected {
					t.Errorf("expected %s, got %s", test.expected, got)
				}
			})
		}
	})

	t.Run("test ints", func(t *testing.T) {
		tests := []struct {
			name     string
			values   []int
			expected int
		}{
			{name: "all set ints", values: []int{1, 2, 3}, expected: 1},
			{name: "last two set ints", values: []int{0, 2, 3}, expected: 2},
			{name: "last one set ints", values: []int{0, 0, 3}, expected: 3},
			{name: "all zero", values: []int{0, 0, 0}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := core.Coalesce(test.values...)
				if got != test.expected {
					t.Errorf("expected %d, got %d", test.expected, got)
				}
			})
		}
	})

	t.Run("test floats", func(t *testing.T) {
		tests := []struct {
			name     string
			values   []float32
			expected float32
		}{
			{name: "all set float32s", values: []float32{1.1, 2.2, 3.3}, expected: 1.1},
			{name: "last two set float32s", values: []float32{0, 2.2, 3.3}, expected: 2.2},
			{name: "last one set float32s", values: []float32{0, 0, 3.3}, expected: 3.3},
			{name: "all zero", values: []float32{0, 0, 0}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				got := core.Coalesce(test.values...)
				if got != test.expected {
					t.Errorf("expected %f, got %f", test.expected, got)
				}
			})
		}
	})
}

func TestIIF(t *testing.T) {
	tests := []struct {
		name      string
		condition bool
		v1        string
		v2        string
		expected  string
	}{
		{condition: true, v1: "1", v2: "2", expected: "1"},
		{condition: false, v1: "1", v2: "2", expected: "2"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := core.IIF(test.condition, test.v1, test.v2)
			if got != test.expected {
				t.Errorf("expecte %s, got %s", test.expected, got)
			}
		})
	}
}
