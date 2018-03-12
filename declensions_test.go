package unit

import "testing"

func TestDeclensions(t *testing.T) {
	tests := []testCase{
		{value: 1, expected: "nominative"},
		{value: 2, expected: "genitive"},
		{value: 3, expected: "genitive"},
		{value: 4, expected: "genitive"},
		{value: 5, expected: "plural"},
		{value: 7, expected: "plural"},
		{value: 10, expected: "plural"},
		{value: 11, expected: "plural"},
		{value: 12, expected: "plural"},
		{value: 15, expected: "plural"},
		{value: 17, expected: "plural"},
		{value: 19, expected: "plural"},
		{value: 21, expected: "nominative"},
		{value: 21.01, expected: "plural"},
		{value: 21.99, expected: "plural"},
		{value: 21.99, expected: "plural"},
		{value: 0.99, expected: "plural"},
		{value: 0.49, expected: "plural"},
		{value: 0.77, expected: "plural"},
	}

	test(t, tests, func(value float64) string { return declensions(value, "nominative", "genitive", "plural") })
}
