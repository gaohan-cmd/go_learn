package groutine

import (
	"fmt"
	"log"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go关键字开启一个goroutine，顺序不一致
		go func(i int) {
			t.Log(i)
		}(i)
		// 当不传入i,那么i则会被多个协程所共享，共享变量会存在竞争条件，所以i这个时候需要锁的机制来完成协程
		// 按上述办法可行 原因是GO的方法调用传递是值传递，14行i的值会被复制一份，在每个协程所拥有变量地址是不一样的，所以不会存在竞争关系
		go func() {
			fmt.Println(i)
		}()
	}

	for i := 0; i < 10; i++ {
		log.Printf("i = %d\n", i)
	}

}
