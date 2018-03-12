package unit

import (
	"fmt"
	"testing"
)

type testCase struct {
	value    float64
	expected string
}

func (test *testCase) Name() string {
	return fmt.Sprintf("Value:'%v',Exprected:'%s'", test.value, test.expected)
}

func TestNanoseconds(t *testing.T) {
	tests := []testCase{
		{value: 1, expected: "1 ns"},
		{value: 10, expected: "10 ns"},
		{value: 100, expected: "100 ns"},
		{value: 1000, expected: "1 µs"},
		{value: 10000, expected: "10 µs"},
		{value: 100000000, expected: "100 ms"},
		{value: 10000000000, expected: "10 sec"},
		{value: 100000000000, expected: "1.67 min"},
		{value: 1000000000000, expected: "16.67 min"},
		{value: 10000000000000, expected: "2.78 hours"},
		{value: 100000000000000, expected: "1.16 days"},
		{value: 1000000000000000, expected: "1.65 weeks"},
		{value: 10000000000000000, expected: "16.53 weeks"},
		{value: 100000000000000000, expected: "3.17 years"},
		{value: 2551, expected: "2.551 µs"},
		{value: 2552, expected: "2.552 µs"},
		{value: 2554, expected: "2.554 µs"},
		{value: 2555, expected: "2.555 µs"},
		{value: 2556, expected: "2.556 µs"},
		{value: 2558, expected: "2.558 µs"},
	}

	test(t, tests, Nanoseconds)
}

func TestMicroseconds(t *testing.T) {
	tests := []testCase{
		{value: 1, expected: "1 µs"},
		{value: 10, expected: "10 µs"},
		{value: 1000, expected: "1 ms"},
		{value: 10000, expected: "10 ms"},
		{value: 10000000, expected: "10 sec"},
		{value: 100000000, expected: "1.67 min"},
		{value: 1000000000, expected: "16.67 min"},
		{value: 10000000000, expected: "2.78 hours"},
		{value: 100000000000, expected: "1.16 days"},
	}

	test(t, tests, Microseconds)
}

func TestMilliseconds(t *testing.T) {
	tests := []testCase{
		{value: 0.0001, expected: "100 ns"},
		{value: 0.1, expected: "100 µs"},
		{value: 1, expected: "1 ms"},
		{value: 10, expected: "10 ms"},
		{value: 1000, expected: "1 sec"},
		{value: 10000, expected: "10 sec"},
		{value: 10000000, expected: "2.78 hours"},
		{value: 100000000, expected: "1.16 days"},
		{value: 1000000000, expected: "1.65 weeks"},
		{value: 10000000000, expected: "16.53 weeks"},
		{value: 100000000000, expected: "3.17 years"},
	}

	test(t, tests, Milliseconds)
}

func TestSeconds(t *testing.T) {
	tests := []testCase{
		{value: 0.0001, expected: "100 µs"},
		{value: 0.1, expected: "100 ms"},
		{value: 1, expected: "1 sec"},
		{value: 10, expected: "10 sec"},
		{value: 1000, expected: "16.67 min"},
		{value: 10000, expected: "2.78 hours"},
		{value: 10000000, expected: "16.53 weeks"},
		{value: 100000000, expected: "3.17 years"},
		{value: 1000000000, expected: "31.71 years"},
	}

	test(t, tests, Seconds)
}

func test(t *testing.T, testCases []testCase, testFunc func(value float64) string) {
	for _, test := range testCases {
		str := testFunc(test.value)
		t.Run(test.Name(), func(tt *testing.T) {
			if str != test.expected {
				tt.Errorf("Expected: %s,Actual: %s", test.expected, str)
			}
		})
	}
}
