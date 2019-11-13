package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr[1] = 2
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3}
	t.Log(len(arr), arr)
	t.Log(len(arr1), arr1)
	t.Log(len(arr2), arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}
}

func TestMutiArray(t *testing.T) {
	arr4 := [...][3]int{{1, 2}, {1, 2, 3}}
	for idx, e := range arr4 {
		t.Log(idx, e)
	}
}
