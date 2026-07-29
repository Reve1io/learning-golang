package tests

import (
	arrays "test/tutorials/Arrays"
	"testing"
)

var synthetics = []struct {
	history1 [][]string
	history2 [][]string
	out      bool
}{
	{
		history1: [][]string{
			[]string{"milk", "bread"},
			[]string{"apple", "banana"},
		},
		history2: [][]string{
			[]string{"milk", "bread"},
			[]string{"apple", "banana"},
		},
		out: true,
	},
	{
		history1: [][]string{
			[]string{"milk", "bread"},
			[]string{"apple", "banana"},
		},
		history2: [][]string{
			[]string{"milk", "bread"},
			[]string{"apple", "banana"},
			[]string{"apple", "banana"},
		},
		out: false,
	},
	{
		history1: [][]string{},
		history2: [][]string{},
		out:      false,
	},
}

func ComparisonOfSlicesTest(t *testing.T) {
	for _, synthetic := range synthetics {
		equal := arrays.AreOrderHistoriesEqual(synthetic.history1, synthetic.history2)

		if equal != synthetic.out {
			t.Fatal("The result does not match the expected result")
		} else {
			t.Log("Ok!")
		}

	}
}
