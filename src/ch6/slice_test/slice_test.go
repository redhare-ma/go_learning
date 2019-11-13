package slice_test

import "testing"

func TestSliceinit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4, 5}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 2)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrouwing(t *testing.T) {
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

//共享存储
func TestSliceShareMemory(t *testing.T) {
	year := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	Q3 := year[5:6]
	t.Log(Q3, len(Q3), cap(Q3))
	Q3[0] = "six"
	t.Log(year, Q2, Q3)

}

func TestSliceEqual(t *testing.T) {
	a := []string{}
	//	b := []int{1, 2, 3, 4}
	t.Log(a, len(a))
	if a == nil {
		t.Log("nil")
	}
}
