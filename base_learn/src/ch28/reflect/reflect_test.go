package reflect

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

// 类型信息（Type）指的是一个对象的具体类型;种类信息（Kind）则是对类型进行更粗略的分类，指的是对象的基本分类
// 举个例子来说，对于一个结构体类型的变量，其类型信息表示这个变量确切的类型是哪个结构体，而种类信息则表示这个变量的种类是结构体
//
//	type Example struct {
//		Name string
//		Age  int
//	}
//
//	func main() {
//		var x Example
//		t := reflect.TypeOf(x)
//		k := t.Kind()
//
//		fmt.Println("类型信息：", t)  // 输出：main.Example
//		fmt.Println("种类信息：", k)  // 输出：struct
//	}
func CheckType(v interface{}) {
	//空接口可以接受所有类型
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Interger")
	default:
		fmt.Println("Unknow", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}

func TestDeepEqual(t *testing.T) {

}

type Employee struct {
	EmployeeID string
	Name       string "format:'normal'" //struct Tag
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	//按照名字获取成员 通过反射的方式 直接返回了名字的值
	t.Logf("Name value(%[1]v),Type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		//返回两个值 ok代表 有没有这个值
		t.Error("Failed to get 'Name' field")
	} else {
		//名字和结构体的对应 都是通过 structTag
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
	t.Log("Update Age:", e)
}

func TestFillNameAndAge(t *testing.T) {
	// 空接口类型(interface{})可以表示任意类型的值。
	setting := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	// 通过反射的方式进行填充
	if err := fillBySettings(&e, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)

}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	// 使用reflect.TypeOf获取参数st的类型，然后调用Kind()方法获取其种类。reflect.Ptr表示指针类型
	// 类型信息描述了具体的类型，而种类信息描述了类型的基本分类。举个例子，一个变量的类型信息可能是 *int，而它的种类信息则是 reflect.Ptr，表示指针类型
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem()方法获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil")
	}
	for k, v := range settings {
		// reflect.ValueOf(st).Elem().Type().FieldByName(k) 通过反射的方式获取结构体的字段
		// reflect.ValueOf(st).Elem().Type().FieldByName(k) 获取结构体的字段的类型
		// reflect.ValueOf(st).Elem().Type().FieldByName(k).Type 获取结构体的字段的类型
		// reflect.ValueOf(st).Elem().Type().FieldByName(k).Type == reflect.TypeOf(v) 判断结构体的字段的类型是否和传入的值的类型一致
		// reflect.ValueOf(st).Elem().FieldByName(k).Set(reflect.ValueOf(v)) 通过反射的方式设置结构体的字段的值
		if field, ok := (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		} else {
			fmt.Printf("k====%v,v====%v\n", k, v)
			fmt.Println("reflect.TypeOf(st)====", reflect.TypeOf(st))
			fmt.Println("reflect.ValueOf(st)====", reflect.ValueOf(st))
			fmt.Println("reflect.ValueOf(st).Elem()====", reflect.ValueOf(st).Elem())
			fmt.Println("reflect.ValueOf(st).Elem().Type()====", reflect.ValueOf(st).Elem().Type())
			// 判断结构体字段的类型是否与传入值的类型一致,结构体字段的类型是通过反射的方式获取的，即根据原始结构体的类型进行匹配
			if field.Type == reflect.TypeOf(v) {
				vstr := reflect.ValueOf(st)
				vstr = vstr.Elem()
				vstr.FieldByName(k).Set(reflect.ValueOf(v))
			}
		}
	}
	return nil

}
