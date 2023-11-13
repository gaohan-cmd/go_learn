package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 协程不安全demo
// 协程并发，导致的协程不安全
func TestCounterThreadUnsafe(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

//解决方式一：
//普通加锁，并加延迟等待协程执行完毕（不推荐）

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				mut.Unlock() // 函数调用完成后 解锁，保证协程安全
			}()
			mut.Lock() // 函数将要调用前，将要计数前，进行加锁，保证协程安全
			counter++
		}()
	}
	time.Sleep(time.Second) // 等待协程执行完毕
	t.Logf("counter= %d", counter)

}

// 解决方式二：
// 推荐！ 使用同步等待队列（WaitGroup）保证顺序执行
// 上述结果正确，但是有一个问题。因为这里有个1秒的延迟等待，保证协程运行完毕再调用结果
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex    //互斥锁
	var wg sync.WaitGroup // 等待队列
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1) // 每次循环都加一个协程
		go func() {
			defer func() {
				mut.Unlock() // 函数调用完成后 解锁，保证协程安全
			}()
			mut.Lock() // 函数将要调用前，将要计数前，进行加锁，保证协程安全
			counter++
			wg.Done() // 每次循环完成后，减一个协程
		}()
	}
	wg.Wait() // 等待协程执行完毕
	t.Logf("counter= %d", counter)

}
