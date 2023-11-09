package map_ext

import "testing"

func TestMapWithFunalue(t *testing.T) {
	// value设置为函数 实现工厂模式
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))

}

// 这里演示其他语言中set集合 使用某个值作为key 数值对应设定为bool
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	// 设置为true是确保数据存入set中 如果设置为false那么按道理应该是不能存入set的 但实际会被存入影响后续if判断是否存在
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = false
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

}
