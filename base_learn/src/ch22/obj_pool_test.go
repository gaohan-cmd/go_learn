package ch22

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {
}

// chan *ReusableObj表示一个能够传递指向ReusableObj对象的指针的通道。这样做的好处是，在将ReusableObj对象放入通道或从通道中取出时，
// 并不需要复制整个对象，而只需要传递指针，从而提高了效率并节省了内存开销
type ObjPool struct {
	bufChan chan *ReusableObj // 用于缓冲可重用对象
}

// 生产指定数量对象的对象池
func NewObjPool(numOfObj int) *ObjPool {
	ObjPool := ObjPool{}
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		ObjPool.bufChan <- &ReusableObj{}
	}
	return &ObjPool
}

// 从对象池中获得对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

// 释放对象池里的对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10) // 生产一个10容量大小的对象池
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil { // 释放一个不存在的对象
	//	t.Error(err)
	//}
	for i := 0; i < 10; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil { // 获取obj
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)                      // 获取成功，答应日志。
			if err := pool.ReleaseObj(v); err != nil { // 释放obj
				t.Error(err)
			}
		}
	}
	fmt.Println("Done.")
}
