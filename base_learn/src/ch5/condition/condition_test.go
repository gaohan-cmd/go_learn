package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	// 判断+a必须为布尔值
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}
