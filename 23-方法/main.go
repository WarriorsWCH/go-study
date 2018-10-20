
package main 

import (
	"fmt"

)
// 方法
// go中的方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），
// 因此自定义类型都可以是方法，而不仅仅是struct
// 方法的声明（定义）
// func (recevier type) methondName(参数列表) (返回值列表){
	//方法体
	//return 返回值 
// }
// 1.参数列表：表示方法输入
// 2.recevier type:表示这个方法和type这个类型进行绑定，或者说该方法作用于type类型
// 3.recevier type:type可以是结构体，也可以是其他的自定义类型
// 4.返回值列表：表示返回值，可以多个
// 5.recevier：就是type类型的一个变量（实例），比如：Person结构体的一个变量（实例）
// 6.方法主题：表示为了实现某一功能代码块
// 7.return语句不是必须的
func main() {
	var p Person
	p.Name = "xiao ming"
	p.test()
	p.say("北京")
	// 方法的调用和传参机制原理：
	// 方法的调用和传参机制和函数基本一样，不一样的地方是方法的调用，会将调用方法的变量，当做实参也传递给方法
	// 变量调用方法时，该变量本身也会作为一个参数传递到方法
	// (如果变量是值类型，则进行拷贝，如果变量是引用类型，则进行地址拷贝)

	var c Circle
	c.radius = 4.0
	fmt.Println(c)
	res := c.area()
	fmt.Println("面积：",res)


	var c2 Circle
	c2.radius = 1.0
	res2 := c2.area2()
	fmt.Println("面积是:",res2)
	fmt.Println("c2 radius=",c2.radius)

	var i integer = 10
	i.print()
	i.change()
	i.print()

	stu := Student{
		Name: "xiao",
		Age: 20,
	}
	fmt.Println(&stu)
}


type Person struct {
	Name string
}
// test方法和Person类型绑定
// test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其他类型变量来调用
// func (p Person) test(){}...p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参相似
func (p Person) test() {
	fmt.Println("test() name=",p.Name)
}
func (p Person) say(city string) {
	fmt.Printf("%v是一个%v市的人\n",p.Name,city)
}



type Circle struct {
	radius float64
}
func (c Circle) area() float64 {

	return 3.14 * c.radius * c.radius
}
// 为了提高效率，通常我们的方法和指针类型绑定
func (c *Circle) area2() float64 {
	c.radius = 10
	return 3.14 * c.radius * c.radius
}

type integer int

func (i integer) print() {
	fmt.Println("i=",i)
}
func (i *integer) change() {
	*i = *i + 1
}

// 方法的注意事项和细节
// 1.结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
// 2.如程序员希望在方法中，修改结构体变量的值，可以通过结构体指针的方式来处理
// 3.go中的方法作用在指定的数据类型上的，因此自定义类型都可以有方法
// 而不仅仅是struct，比如：int，float64等都可以有方法
// 4.方法的访问范围控制规则，和函数一样，方法名首字母小写，只能在本包访问，方法首字母大写，可以在本包和其他包访问
// 5.如果一个类型实现了String()这个方法，namefmt.Println默认会调用这个变量的String()进行输出
type Student struct {
	Name string
	Age int
}
func (s *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]",s.Name,s.Age)
	return str
}















