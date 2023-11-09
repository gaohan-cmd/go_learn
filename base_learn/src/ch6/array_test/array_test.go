package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	var aff [2]int = [...]int{1, 2}
	t.Log(aff)
	t.Log(arr)
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr2); i++ {
		t.Log(arr2[i])
	}

	for idx, e := range arr2 {
		t.Log(idx, e)
	}

	for _, e := range arr2 {
		t.Log(e)
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	// 负数不允许
	arr_3 := arr[0:3]
	t.Log(arr_3)
}
