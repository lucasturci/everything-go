package comparator

import (
	"testing"
)

func TestLessComparator(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected bool
	}{
		{"less than", 1, 2, true},
		{"equal", 2, 2, false},
		{"greater than", 3, 2, false},
	}

	less := Less[int]{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := less.Less(tt.a, tt.b); got != tt.expected {
				t.Errorf("Less.Less(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestGreaterComparator(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{"less than", "a", "b", false},
		{"equal", "b", "b", false},
		{"greater than", "c", "b", true},
	}

	greater := Greater[string]{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greater.Less(tt.a, tt.b); got != tt.expected {
				t.Errorf("Greater.Less(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

func TestCustomComparator(t *testing.T) {
	// Test valid custom comparator
	t.Run("valid custom comparator", func(t *testing.T) {
		customLess := func(a, b int) bool {
			return a%2 < b%2 // Compare by remainder when divided by 2
		}
		comp := Custom(customLess)

		tests := []struct {
			a, b     int
			expected bool
		}{
			{1, 2, false}, // 1%2 = 1, 2%2 = 0
			{2, 3, true},  // 2%2 = 0, 3%2 = 1
			{4, 6, false}, // 4%2 = 0, 6%2 = 0
		}

		for _, tt := range tests {
			if got := comp.Less(tt.a, tt.b); got != tt.expected {
				t.Errorf("Custom.Less(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		}
	})

	// Test panic with nil custom comparator
	t.Run("nil custom comparator", func(t *testing.T) {
		comp := CustomComparator[int]{}
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic with nil custom comparator, but it didn't panic")
			}
		}()
		comp.Less(1, 2)
	})
}
