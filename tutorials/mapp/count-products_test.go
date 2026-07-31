package mapp_test

import (
	"maps"
	"test/tutorials/mapp"
	"testing"
)

func TestCountProducts(t *testing.T) {
	tests := []struct {
		name  string
		input [][]string
		want  map[string]int
	}{
		{
			name: "1. default case",
			input: [][]string{
				{"milk", "strawberry"},
				{"apple", "banana"},
			},
			want: map[string]int{"apple": 1, "banana": 1, "milk": 1, "strawberry": 1},
		},
		{
			name: "2. repeat value case",
			input: [][]string{
				{"milk", "banana"},
				{"apple", "banana"},
			},
			want: map[string]int{"apple": 1, "banana": 2, "milk": 1},
		},
		{
			name: "3. Exists default value case",
			input: [][]string{
				{"milk", "strawberry"},
				{"", "banana"},
			},
			want: map[string]int{"": 1, "banana": 1, "milk": 1, "strawberry": 1},
		},
		{
			name:  "4. Slice is nil case",
			input: nil,
			want:  map[string]int{},
		},
		{
			name:  "5. Slice is empty case",
			input: [][]string{},
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapp.CountProducts(tt.input)
			equal := maps.Equal(got, tt.want)
			if got == nil {
				t.Errorf("CountProducts(%v) = %v; want %v", tt.input, got, tt.want)
			}
			if equal != true {
				t.Errorf("CountProducts(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
