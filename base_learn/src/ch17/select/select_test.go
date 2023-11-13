package _select

import (
	"fmt"
	"testing"
	"time"
)

// 多路选择和超时控制
func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService(): // 从管道中读取结果，这里的case和下面的case没有先后顺序，哪个先执行完毕就先执行哪个
		t.Log(ret)
	case <-time.After(time.Millisecond * 100): // 超时控制，如果管道中没有数据(等待100ms），那么就会执行这个case
		t.Error("time out")
	}

}
