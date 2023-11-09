package map_test

import (
	"testing"
)

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	m2 := map[int]int{}
	m2[4] = 16
	t.Log("len: ", len(m2))
	m3 := make(map[int]int)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	// 对于Go语言中 默认初始化为0值（避免了空指针异常） 那区分赋值的0和初始化的0 按如下办法
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not exist")
	}
}

func TestTraveMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	// map 这里返回是key和value  array 返回是索引和数值 要进行区分
	for k, v := range m1 {
		t.Logf("key %d ,value %d", k, v)
	}

}
