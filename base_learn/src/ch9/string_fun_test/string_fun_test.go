package string_fun_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestStrngFn(t *testing.T) {
	s := "A,B,C,D,E,F,G,H,I"
	parts := strings.Split(s, ",")
	for _, i := range parts {
		t.Logf("testing %s", i)
	}
	t.Log(strings.Join(parts, "-"))

}

func TestConv(t *testing.T) {
	// 整形转换成字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)

	// shiyong
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
