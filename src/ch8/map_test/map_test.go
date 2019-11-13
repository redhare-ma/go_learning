package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[2] = true
	n := 2
	t.Log(mySet[n])
	if mySet[n] {
		t.Logf("%d is exting", n)
	} else {
		t.Logf("%d is not exting", n)
	}

	mySet[1] = true
	t.Logf("mySet len id %d", len(mySet))

	delete(mySet, 2)
	if mySet[n] {
		t.Logf("%d is exting", n)
	} else {
		t.Logf("%d is not exting", n)
	}

	mySet[1] = true
	t.Logf("mySet len id %d", len(mySet))
}
