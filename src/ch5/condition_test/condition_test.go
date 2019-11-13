package condition_test

import "testing"

//func TestIfMultiSec(t *testing.T) {
//	if v, err := someFun(); err == nil {
//		t.Log("1 == 1")
//	} else {
//		t.Log("1 == 1")
//	}
//}

func TestSwitchMulticCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is nopt 0-3")
		}
	}
}

func TestSwitchMulticCase1(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("it is nopt 0-3")
		}
	}
}
