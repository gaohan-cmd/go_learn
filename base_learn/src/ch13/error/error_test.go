package error

import (
	"errors"
	"fmt"
	"testing"
)

// 区分错误类型
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 0 || n > 100 {
	//	return nil, errors.New("n should be in [2,100]")
	//}
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}
	fibList := []int{1, 1}
	// fibList是要追加元素的切片，fibList[i-2]+fibList[i-1]是要追加的一个或多个元素。append函数会将元素追加到slice的末尾，并返回一个新的切片
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	// 这里需要对Go语言进行校验
	t.Log(GetFibonacci(-10))
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		if err == LargerThanHundredError {
			fmt.Println("It is larger.")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
