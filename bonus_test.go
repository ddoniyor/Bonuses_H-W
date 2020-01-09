package main

import "testing"

func Test_sumOfBonuses(t *testing.T) {

	tests := []struct {
		name  string
		sales []int
		want  int
	}{
		// TODO: Add test cases.

		// TODO: Add test cases.
		{"Correct sum of all Bonuses", []int{12_000, 15_000}, 350},
		{"Without Bonuses", []int{8_000, 8_000}, 0},
		{"Bound line", []int{10_000, 10_000}, 0},
	}
	for _, test := range tests {
		got := sumOfBonuses(test.sales)
		if got != test.want {
			t.Error("Bonus test:", test.name, "got:", got, "want:", test.want)
		}
	}
}
