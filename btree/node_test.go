package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	// 	   | 12 | 32 | 44 | 60 |
	// 	| 10 | 13 | 37 | 57 | 72 |
	tt := []struct {
		name     string
		input    int
		expected Node
	}{
		{
			name:  "find root key 12",
			input: 12,
			expected: Node{
				key:   []int{12, 32, 44, 60},
				child: fixtureChildNodes(),
			},
		},
		{
			name:  "find root key 32",
			input: 32,
			expected: Node{
				key:   []int{12, 32, 44, 60},
				child: fixtureChildNodes(),
			},
		},
		{
			name:  "find root key 44",
			input: 44,
			expected: Node{
				key:   []int{12, 32, 44, 60},
				child: fixtureChildNodes(),
			},
		},
		{
			name:  "find root key 60",
			input: 60,
			expected: Node{
				key:   []int{12, 32, 44, 60},
				child: fixtureChildNodes(),
			},
		},
		{
			name:  "find child key 10",
			input: 10,
			expected: Node{
				key:   []int{10},
				child: []Node{},
			},
		},
		{
			name:  "find child key 13",
			input: 13,
			expected: Node{
				key:   []int{13},
				child: []Node{},
			},
		},
		{
			name:  "find child key 37",
			input: 37,
			expected: Node{
				key:   []int{37},
				child: []Node{},
			},
		},
		{
			name:  "find child key 57",
			input: 57,
			expected: Node{
				key:   []int{57},
				child: []Node{},
			},
		},
		{
			name:  "find child key 72",
			input: 72,
			expected: Node{
				key:   []int{72},
				child: []Node{},
			},
		},
		{
			name:     "don't find child key 9",
			input:    9,
			expected: Node{},
		},
		{
			name:     "don't find child key 11",
			input:    11,
			expected: Node{},
		},
		{
			name:     "don't find child key 20",
			input:    20,
			expected: Node{},
		},
		{
			name:     "don't find child key 36",
			input:    36,
			expected: Node{},
		},
		{
			name:     "don't find child key 40",
			input:    40,
			expected: Node{},
		},
		{
			name:     "don't find child key 48",
			input:    48,
			expected: Node{},
		},
		{
			name:     "don't find child key 59",
			input:    59,
			expected: Node{},
		},
		{
			name:     "don't find child key 65",
			input:    65,
			expected: Node{},
		},
		{
			name:     "don't find child key 74",
			input:    74,
			expected: Node{},
		},
	}

	root := NewNode()
	root.key = append(root.key, 12, 32, 44, 60)
	root.child = append(root.child, fixtureChildNodes()...)

	for _, tc := range tt {
		t.Run(fmt.Sprintf("recursive %s", tc.name), func(t *testing.T) {
			node := root.SearchRecursive(tc.input)

			assert.Equal(t, tc.expected, node)
		})

		t.Run(fmt.Sprintf("in place %s", tc.name), func(t *testing.T) {
			node := root.SearchInPlace(tc.input)

			assert.Equal(t, tc.expected, node)
		})
	}
}

func fixtureChildNodes() []Node {
	return []Node{
		{
			key:   []int{10},
			child: []Node{},
		},
		{
			key:   []int{13},
			child: []Node{},
		},
		{
			key:   []int{37},
			child: []Node{},
		},
		{
			key:   []int{57},
			child: []Node{},
		},
		{
			key:   []int{72},
			child: []Node{},
		},
	}
}
