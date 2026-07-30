package arrays_test

import (
	arrays "test/tutorials/Arrays"
	"testing"
)

var synthetics = []struct {
	testName string
	history1 [][]string
	history2 [][]string
	out      bool
}{
	{
		testName: "1. Сonverges",
		history1: [][]string{
			{"milk", "bread"},
			{"apple", "banana"},
		},
		history2: [][]string{
			{"milk", "bread"},
			{"apple", "banana"},
		},
		out: true,
	},
	{
		testName: "2. Varying lengths",
		history1: [][]string{
			{"milk", "bread"},
			{"apple", "banana"},
		},
		history2: [][]string{
			{"milk", "bread"},
			{"apple", "banana"},
			{"apple", "banana"},
		},
		out: false,
	},
	{
		testName: "3. Equal empty string values",
		history1: [][]string{{""}},
		history2: [][]string{{""}},
		out:      true,
	},
	{
		testName: "4. Arrays is nil",
		history1: nil,
		history2: nil,
		out:      true,
	},
}

func TestAreOrderHistoriesEqual(t *testing.T) {
	for _, synthetic := range synthetics {
		t.Run(synthetic.testName, func(t *testing.T) {
			equal := arrays.AreOrderHistoriesEqual(synthetic.history1, synthetic.history2)
			if equal != synthetic.out {
				t.Errorf("The result does not match the expected result:\n expect - %v\n fact - %v\n", synthetic.out, equal)
			}
		})
	}
}
