package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractMathExpression(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "1 + 1 and 1 + 2 is not equal",
			expected: []string{
				"1 + 1",
				"1 + 2",
			},
		},
	}
	for _, c := range cases {
		assert.ElementsMatch(t, c.expected,extractMathExpressions(c.input))
	}
}