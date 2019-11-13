package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time Spent:", time.Since(start).Seconds())
		fmt.Println(ret, n, start)
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * time.Duration(op))
	fmt.Println("time.Second: ", time.Second)
	return op
}

func TestFun(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}
