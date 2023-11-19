package benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkConcatStringByAdd(b *testing.B) {
	//与性能测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//测试代码
		//b.N 自动处理 知道持续的时间足够长
	}
	b.StopTimer()
	//与性能测试无关的代码
}

// 目的 处理一个切片中字符串的整合工作的效率比较
func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("12345", ret)
}

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	elems := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	assert.Equal("12345", buf.String())
}
