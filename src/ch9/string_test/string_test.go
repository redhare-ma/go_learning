package string_test

import "testing"

func TestStringInit(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5" //可以存储任何类型的二进制数据
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s, len(s))
	//s[1] = '3' //string 是不可变的byte slice
	//t.Log(s)

	s = "中"
	t.Log(len(s)) //是byte数

	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x %[1]d", c)
	}
}
