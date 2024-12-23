package main

import "testing"

func TestSolution(t *testing.T) {
	input := []byte(`kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`)

	t.Run("part 1", func(t *testing.T) {
		expected := 7
		if got := partOne(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})

	t.Run("part 2", func(t *testing.T) {
		expected := "co,de,ka,ta"
		if got := partTwo(input); got != expected {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}
