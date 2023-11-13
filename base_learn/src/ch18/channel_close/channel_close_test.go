package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 关闭管道--解决办法二
		close(ch)
		//  关闭后继续发送消息 会引发panic
		// ch <- 11
		// 等待协程执行完毕
		wg.Done()
	}()
}

// 这里接收者和产生着 需要注意的是数量保持一致 因为每个接收者都需要知道什么时候关闭管道（即收到关闭信号）
// 解决方法一：设置一个结束变量（比如-1） 当接收到-1时候生产者生成完成了，接收者可以退出了。但是
// 这也会引发一个新的问题 当放置一个-1时候会被多个receive中某个接收到 如果有多个receiver那就必须知道有几个reiver 需要放几个-1(这个方法写那么耦合度太高，不推荐）
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			// 从管道中读取数据
			data := <-ch
			//fmt.Printf("data = %d\n", data)
			fmt.Println(data)
		}
		wg.Done()
	}()
}

func dataReceiver2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		// 正常写法 通过ok来判断管道是否关闭，需要判断
		for {
			// 从管道中读取数据 true是正常接收状态 false状态是通道关闭，channel关闭后，如果没有拿ok去接收，再取值就是对应类型的零值
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}

		//for i := 0; i < 11; i++ {
		//	// 在这里因为只有10个数 所以第十一个数字发送接收不到 在这里测试不进行判断
		//	// 这里结果显示 程序不会阻塞在这里 因为通道已经被关闭了 所以会返回对应类型的零值
		//	data := <-ch
		//	fmt.Println(data)
		//}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver2(ch, &wg)
	// wg.Wait(): 这个方法用于阻塞程序的执行，直到WaitGroup内部的计数归零
	wg.Wait()

}
