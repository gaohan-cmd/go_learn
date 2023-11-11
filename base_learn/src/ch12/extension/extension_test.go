package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 匿名嵌套类型 然后可以删掉func (d *Dog) Speak() 和func (d *Dog) SpeakTo(host string) 方法
type Dog struct {
	Pet
}

// 这中写法调用时候需要自己重写 `SpeakTo` 方法
/*type Dog struct {
	p *Pet
}*/

/*func (d *Dog) Speak() {
	fmt.Println("wang!")
	d.p.Speak()
}


func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}*/

func (d *Dog) Speak() {
	// 当dog作为匿名嵌套类型时，依然是不能支持方法的重载的，只能调用原本父类方法
	fmt.Println("wang!")
}

func TestDog(t *testing.T) {
	// go不支持显式类型转换 即不支持下述写法
	// var dog Pet = new(Dog)
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("Chao")
}
