package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m[2])
	t.Logf("len m=%d", len(m))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[1], m1[2])
	if v, ok := m1[2]; ok {
		t.Logf("key 2's values is %d", v)
	} else {
		t.Log("key 2's values not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
