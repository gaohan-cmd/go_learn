package obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

func TestSynPool(t *testing.T) {
	// 用于存储临时对象的共享池,sync.Pool对象的New字段指定的匿名函数会在需要新对象时被调用，输出一条日志并返回值100作为新的对象
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 100
		},
	}

	v := pool.Get().(int) //断言 (int)：表示将接口类型的值断言为int类型
	fmt.Println(v)
	pool.Put(3)
	//runtime.GC() 不推荐人为触发  GC会清除sync.pool中缓存的对象
	v1, _ := pool.Get().(int) //这里继续进行断言 由于私有对象已经存在，所以不会再次调用New函数，而是直接返回之前缓存的对象
	fmt.Println(v1)

	//不Put就不会进入sync.Pool
}

// 测试多线程环境下对象池的缓存机制
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(122)
	pool.Put(123)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
