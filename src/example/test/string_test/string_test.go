package string_test

import "testing"

func TestString(t *testing.T) {

	var s string
	t.Log(s, len(s))
	s = "hello"
	t.Log(s, len(s))

	//s[1] = '3' string是不可变的byte slice

	s = "你好"
	t.Log(s, len(s)) //byte数

	c := []rune(s) //rune 表示string的unicode
	t.Log(c)
}

func TestStringToRune(t *testing.T) {
	s := "hello world"
	for _, c := range s {
		t.Logf("%[1]c %[1]d",c)
	}
}
