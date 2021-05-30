package basic

import (
	"testing"
)

func Test1(t *testing.T) {
	var a int = 100
	var b int = 200

	var (
		c int = 300
		d int = 400
	)
	t.Log(a, b)
	a, b = b, a
	t.Log(a, b)

	t.Log(c, d)

	f := "golang"
	t.Log(f)

	const (
		day1 = iota + 1
		day2
		day3
	)

	t.Log(day1, day2, day3)
}

func TestFib(t *testing.T) {
	var (
		pre  int = 1
		next int = 1
	)
	for i := 0; i < 10; i++ {
		t.Log(next)
		tmp := pre
		pre = next
		next = tmp + pre
	}
}

// 1 1 2 3 5 8
