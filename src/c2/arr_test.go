package basic

import "testing"

func TestArr1(t *testing.T) {
	var arr1 [5]int
	arr1[1] = 100
	t.Log(arr1)

	// var arr2 = [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	t.Log(arr2)

	arr3 := [2][2]int{{1, 2}, {3, 4}}
	t.Log(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(arr4)

	for i := 0; i < len(arr4); i++ {
		t.Logf("byLen  %d", arr4[i])
	}

	for i, e := range arr4 {
		t.Logf("index:%d,element:%d", i, e)
	}

	t.Log(arr4[1:4])
	t.Log(arr4[:len(arr4)-1])
	t.Log(arr4[3:])

}
