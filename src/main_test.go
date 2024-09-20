package main

import "testing"

func Test_Add_Should_Add_Given_Args(t *testing.T) {
	cases := []struct {
		args     []int
		expected int
	}{
		{
			args:     []int{1, 2, 3},
			expected: 6,
		},
		{
			args:     []int{0, 0, 0},
			expected: 0,
		},
		{
			args:     []int{1, -1},
			expected: 0,
		},
		{
			args:     []int{},
			expected: 0,
		},
	}

	for _, v := range cases {
		res := add(v.args...)
		if res != v.expected {
			t.Errorf("expected %d, got %d", v.expected, res)
		}
	}
}
