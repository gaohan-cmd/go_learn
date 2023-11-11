package polymorphism

import "testing"

// 给string起一个别名
type Code string
type Programmer2 interface {
	WriteHelloWorld() Code
}
type JavaProgrammer struct {
}

// 实现多态 同签名接口
func (p *JavaProgrammer) WriteHelloWorld() Code {

	return "system.out.println(\"Hello World.\")"
}

func TestPolymorphism(t *testing.T) {
	// 下述两种写法都可以
	//javaProgrammer:=&JavaProgrammer{}
	javaProgrammer := new(JavaProgrammer)
	javaProgrammer.WriteHelloWorld()

}
