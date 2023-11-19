package deepequal

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	Employee string
	Name     string
	Age      int
}
type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// 切片和map只能和空集进行比较
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}

	//t.Log(a == b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{2, 4, 6}
	s3 := []int{1, 4, 6}
	//t.Log(s1 == s2)
	t.Log(reflect.DeepEqual(s1, s2))
	t.Log(reflect.DeepEqual(s2, s3))
}

// 关于反射的要求
// 提高了程序的灵活性 降低了程序的可读性 降低了程序的性能
// 构建一个 通用的填充方法
// 公共方法解决不同结构体的填充
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	//func (v value) Elem() Value
	//Elem returns the value that the interface{} v contains or that the pointer
	//It panics if v's Kind is not
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		//elem()获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			//判断是不是结构

			//如果不满足以上条件 不是指针也不是结构 就会返回
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	//传入的是一个指针类型 FieldByName只能从结构类型上去获得
	//从指针类型到指针类型指向的结构，需要使用一个Elem()这样的方法处理 获得了指针指向的结构
	//便利存入的Map k代表类型名字 观测本程序当中是否是Field的名字
	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		//Map和field的类型一致
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			//满足条件后 就设置 Value部 分
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

//和配置有关的程序 需要灵活性的程序 我们都可以使用反射完成
//在性能特别高的情况下 我们要注意 可读性疯狂下降
