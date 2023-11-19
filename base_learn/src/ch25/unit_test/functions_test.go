package unit_test_test

import (
	"ch25/unit_test"
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	//表格测试法
	//一组输入的组合 对一组组合期待的数值
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret := unit_test.Square(inputs[i])
		if ret != expected[i] { //期待的结果和实际的结果不一致
			t.Errorf("input is %d,the expected is %d,the actual %d", inputs[i], expected[i], ret)
		}
	}
}

//内置单元测试框架
//Fail,Error 该测试失败，该测试继续，其他测试继续执行
//FailNow,Fatal 该测试失败，该测试中止,其他测试继续执行

func TestErrorIncode(t *testing.T) {
	fmt.Println("Start")
	t.Error("error")
	fmt.Println("End")
}

func TestFailIncode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("error")
	fmt.Println("End")
}

//代码覆盖率
//coverage  go test -v -cover
