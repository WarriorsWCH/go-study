
package main 

import (
	"fmt"
)


// 接口
// 在go中多态特性主要通过接口体现的
// interface类型可以定义一组方法，但是这些不需要实现，并且interface不能包含任何变量
// 到某个自定义类型（比如结构体）要使用的时候，再根据具体情况把这些方法写出来（实现）
// type 接口名 interface {
	// method1(参数列表)返回值列表
	// mothod2(参数列表)返回值列表
// }
// func (t 自定义类型) method1(参数列表)(返回值列表){}
// func (t 自定义类型) method2(参数列表)(返回值列表){}
type AInterface interface {
	BInterface
	CInterface
	Say()
}
type BInterface interface {
	Hello()
}
type CInterface interface {
	C()
}

type Student struct {
	Name string
}
 
func (s Student) Say() {
	fmt.Println("student say")
}

func (s Student) Hello() {
	fmt.Println("student hello")
}

func (s Student) C() {

}


type integer int 

func (i integer) Say(){
	fmt.Println("integer say",i)
}

func (i integer) Hello(){
	fmt.Println("integer hello",i)
}

func (i integer) C(){
	fmt.Println("C")
}
func main(){
	// 1.接口里的所有方法没有方法体，即接口的方法都是没有实现的方法
	// 接口体现了程序设计的多态和高内聚低耦合的思想
	// 2.go中的接口，不需要显示的实现，只要一个变量，含有接口类型中的所有方法，
	// name这个变量就实现这个接口，因此，go中没有implement这样的关键字
	
	// 接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	var stu Student
	var a AInterface = stu 
	fmt.Printf("a=%v,stu=%v\n",&a,&stu)
	// a=0xc42000e1e0,stu=&{}
	a.Say()
	var s BInterface = stu
	s.Hello()


	// 接口中所有的方法都没有方法体，即都是没有实现的方法
	// 在go中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
	// 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例（变量）赋给接口类型
	// 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

	var i integer = 10
	var b AInterface = i
	b.Say()

	// 一个自定义类型可以实现多个接口
	var stu2 Student
	var a2 AInterface = stu2
	var b2 BInterface = stu2

	a2.Say()
	b2.Hello()

	// go接口中不能有任何变量
	// 一个接口（比如 A接口）可以继承多个别的接口（比如 B,C）
	// 这时如果要实现 A接口，也必须将B,C接口的方法也全部实现

	var t T = stu
	fmt.Println(t)
	// t.Say()//t.Say undefined (type T is interface with no methods)

	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2,t)//8.8 8.8
}


type T interface{}










