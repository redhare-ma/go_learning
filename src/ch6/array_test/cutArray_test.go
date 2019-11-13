package array_test

import "testing"

func TestCutArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	t.Log(arr[1:2])
	t.Log(arr[1:3])
	t.Log(arr[1:len(arr)])
	t.Log(arr[2:])
	t.Log(arr[:2])
	t.Log(arr[:])
	//不能为负数
	//t.Log(arr[:-1])
}
