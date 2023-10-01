package main

import "testing"

type testRun struct {
	name    string
	numbers []int
	target  int
	want    result
}

func TestCheckNumbers(t *testing.T) {
	var tests = []testRun{
		{"0 and 3", []int{10, 15, 3, 7}, 17, result{true, 0, 3}},
		{"0 and 3", []int{10, 15, 3, 8}, 17, result{false, -1, -1}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := CheckForSum(tt.numbers, tt.target)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
