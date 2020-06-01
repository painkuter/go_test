package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateBracket(t *testing.T) {
	var m = map[rune]rune{'(': ')', '{': '}'}
	var cases = map[string]bool{
		"Course sir ( people worthy horses { add entire ) suffer }. Painful so he an { comfort ( is ) manners } . How one dull get busy ) dare ( far. To things so denied admire. ": false,
		"({}){)":           false,
		"()()(){}{}":       true,
		"(((({}{}{}{}))))": true,
	}

	f := func(str []rune) bool {
		var stack []rune
		for _, v := range str {
			switch v {
			case '(', '{':
				{
					stack = append(stack, v)
				}
			case ')', '}':
				{
					if len(stack) == 0 {
						return false
					}
					if m[v] == stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
					}
				}

			default:
			}
		}
		return len(stack) == 0
	}
	for testCase, testResult := range cases {
		assert.Equal(t, testResult, f([]rune(testCase)))
	}
}
