package refl

import (
	"fmt"
	"reflect"
	"unsafe"
)

type X int

//type和kind type表示当前类型  kind表示基础类型
func Test1() {
	var a X = 100

	t := reflect.TypeOf(a)

	fmt.Println(t.Name(), t.Kind())
}

//测试一下循环读取反射中的属性
type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func Test2() {
	var m manager

	t := reflect.TypeOf(m)

	//按照名称
	name, _ := t.FieldByName("name")
	fmt.Println(name.Type, name.Name)

	//按照索引支持多级
	title := t.FieldByIndex([]int{1})
	fmt.Println(title.Name, title.Type)
}

//查看结构体方法（函数如果是小写的开头的话 NumMethod是0）
type A int

type B struct {
	A
}

func (a A) Av()  {}
func (a *A) Ap() {}

func (b B) Bv()  {}
func (b *B) Bp() {}

func Test3() {
	var b B

	t := reflect.TypeOf(&b)
	s := []reflect.Type{t, t.Elem()}

	for _, t := range s {

		fmt.Println(t, ":")

		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(" ", t.Method(i))
		}
	}

}

//遍历类型属性和标签（标签在我们实际生产中用到的机会很多）
type user4 struct {
	name string `field:"name" type:"varchar(50)"`
	age  int    `field:"age" type:"int"`
}

func Test4() {
	var u user4

	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%s:%s %s\n", f.Name, f.Tag.Get("field"), f.Tag.Get("type"))
	}
}

func Test5() {
	a := 100

	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va.CanAddr(), va.CanSet())
	fmt.Println(vp.CanAddr(), vp.CanSet())
	fmt.Println(vp)
}

//验证类型转换
func Test6() {
	/*c := make(chan int,4)
	v := reflect.ValueOf(c)

	if v.TrySend(reflect.ValueOf(100)){
		fmt.Println(v.TryRecv())
	}*/

	var b interface{} = (*int)(nil)
	iface := (*[2]uintptr)(unsafe.Pointer(&b))

	fmt.Println(iface, iface[1] == 0)
}
