package singletion

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

//go如何完成常见的并发任务 在Go语言中如何实现这样的并发任务
//仅仅执行一次  多线程环境下，某段代码，只执行一次
//有方 法保证里面的函数在多线程环境下只会被运行一次

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	//确保在多线程环境下只执行一次
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}
func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
