package interfacetest

import "testing"

// 接口定义,接口是一组方法签名的组合，通过接口可以实现多态
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

// 当有具体的实现类去实现接口时，就可以调用接口的方法，这里需要方法名（签名）一致
// 这里没有去依赖上面的Programmer方法 只是保证签名一样
func (p *GoProgrammer) WriteHelloWorld() string {

	return "fmt.Println(\"Hello World.\")"
}

func TestClient(t *testing.T) {

	var p Programmer // 接口变量
	p = new(GoProgrammer)

	//p2:=GoProgrammer{}
	t.Log(p.WriteHelloWorld())
}
