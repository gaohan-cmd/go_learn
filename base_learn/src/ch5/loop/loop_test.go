package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 10 {
		t.Log(n)
		n++
	}
}
