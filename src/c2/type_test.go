package basic

import (
	"testing"
)

func TestArr(t *testing.T) {
	var arr = [...]int{1, 2, 3, 4, 5}
	var arr1 = [...]int{1, 2, 3, 4, 6}
	var arr2 = [...]int{1, 2, 3, 3, 5}
	var arr3 = [...]int{1, 2, 3, 4, 5}

	t.Log(arr == arr1)
	t.Log(arr == arr2)
	t.Log(arr == arr3)
}

type MyInt32 int32

func TestTypes(t *testing.T) {
	var a int16 = 1
	var b int32 = 2
	var c MyInt32 = 1

	b = int32(2)

	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	var a = "golang"
	var aPtr = &a
	t.Log(a, aPtr)
	t.Logf("%T  %T", a, aPtr)

	var b = [...]string{"halo", "golang"}
	var bPtr = &b
	t.Log(b, bPtr)
	t.Logf("%T   %T", b, bPtr)
}

func TestInitVal(t *testing.T) {
	var a string
	var c int
	var d bool

	t.Log(a, c, d)
}

func TestSwitch(t *testing.T) {
	var a = "2"

	switch a {
	case "1", "2":
		t.Log("1", "2")
	case "3", "4":
		t.Log("3", "4")
	}

	for i := 0; i < 10; i++ {
		t.Log(i % 2)
		if i%2 == 0 {
			t.Log("even")
		} else if i%2 == 1 {
			t.Log("odd")
		} else {
			t.Log("unknow")
		}
		// switch i {
		// case i%2 == 0:
		// 	t.Log("even")
		// case i%2 == 1:
		// 	t.Log("odd")
		// default:
		// 	t.Log("unknow")
		// }
	}

	var num = 0
	for num < 500 {
		t.Log(num)
		num++
	}

	var count = 0
	for {
		count++
		t.Logf("count:%d", count)
		if count >= 100 {
			break
		}
	}
}

func TestCondition(t *testing.T) {
	var a = 0
	var b = 0
	if b = a + 1; b > 0 {
		t.Log("ok")
	} else {
		t.Log("...")
	}
}
