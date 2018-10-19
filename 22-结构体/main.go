


package main 

import(
	"fmt"
	"encoding/json"
)


// 结构体
// go也支持面向对象编程、
// 但是和传统的面向对象编程有区别，并不是纯粹的面向对象语言，所以我们说go支持面向对象编程特性是比较准确的
// go没有类，go语言的结构体和其他编程语言的类有同等地位
// 你可以理解go是基于struct来实现oop特性的
// go面向对象编程非常简洁，去掉了传统oop语言的继承、方法重载、构造函数和析构函数、隐藏this指针等等
// go仍然有面向对象编程的继承，封装和多态的特性，只是实现方式和其他oo语言不同
// go的继承是通过匿名字段来实现的
// go的面向对象很优雅，oop本身就是语言类型系统的一部分，通过接口interface关联，耦合性地，也非常灵活
// 面向接口编程是非常重要的特性


type Cat struct {
	Name string
	Age int
	Color string
	Hobby string
}

type Person struct {
	Name string
	Age int
}


func main(){
	// 创建结构体变量和访问结构体字段
	// 方式一
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Color = "白色"
	cat1.Hobby = "玩"
	fmt.Println(cat1)

	// 方式二
	cat2 := Cat{"小黑",22,"黑色","睡"}
	fmt.Println(cat2)

	// 方式三
	var cat3 *Cat = new(Cat)
	cat3.Name = "小红"
	(*cat3).Age = 33
	cat3.Color = "红色"
	(*cat3).Hobby = "吃"
	fmt.Printf("cat3=%p cat3=%v *cat3=%v\n",cat3,cat3,*cat3)
	// cat3=0xc420054140 cat3=&{小红 33 红色 吃} *cat3={小红 33 红色 吃}


	// 方式四
	var cat4 *Cat = &Cat{}
	(*cat4).Name = "小绿"
	cat4.Age = 55
	cat4.Color = "绿色"
	cat4.Hobby = "笑"
	fmt.Println(*cat4)
	// 第三种和第四种方式返回的是结构体指针
	// 结构体指针访问字段的标准方式是：(*结构体指针).字段名
	// 但go做了一个简化，也支持 结构体.字段名，更加符合程序员的使用习惯
	// go编译器底层对person.Name做了简化(*person).Name


	var p1 Person
	p1.Age = 22
	p1.Name = "jack"

	var p2 *Person = &p1
	fmt.Println((*p2).Age, p2.Age, p1.Age)
	p2.Name = "Tom"
	fmt.Println(p1.Name)
	fmt.Printf("p1的地址%p，p2的地址%p，p2的值%p\n",&p1,&p2,p2)
	// p1的地址0xc42000a060，p2的地址0xc42000c030，p2的值0xc42000a060


	// 结构体的所有字段在内存中是连续的
	r1 := Reat1{Point{1,2}, Point{3,4}}
	fmt.Printf("r1.leftUp.x的地址=%p\n",&r1.leftUp.x)//0xc420018200
	fmt.Printf("r1.rightDown.x的地址=%p\n",&r1.rightDown.x)//0xc420018210
	fmt.Printf("r1.leftUp.y的地址=%p\n",&r1.leftUp.y)//0xc420018208
	fmt.Printf("r1.rightDown.y的地址=%p\n",&r1.rightDown.y)//0xc420018218

	// r2有两个*Point类型，这两个*Point类型本身的地址是连续的，但是他们指向的地址不一定是连续的
	r2 := Reat2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.leftUp=%p,&r2.leftUp=%p, *r2.leftUp=%v,x=%v\n",r2.leftUp,&r2.leftUp,*r2.leftUp,r2.leftUp.x)
	// r2.leftUp=0xc4200160f0,&r2.leftUp=0xc42000e210, *r2.leftUp={10 20},x=10

	var a A
	var b B
	// 可以转换，但是有要求，就是结构体的字段要求完全一样，名字、个数和类型
	a = A(b)
	fmt.Println(a,b)


	type intger int
	var i intger = 10
	var j int = 20
	j = int(j)//不可以使用j = i
	fmt.Println(i,j)


	// 结构体进行type重新定义（相当于取别名），go认为是新的数据类型，但是相互间可以强转
	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)
	fmt.Println(stu1,stu2)

	// struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常用的使用场景就是序列化和反序列化
	monster := Monster{"牛魔王",500,"芭蕉扇"}
	// 将monster序列化为json字符串
	jsonStr, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("处理错误")
	}
	fmt.Println("jsonStr",string(jsonStr))
}


type Point struct {
	x int
	y int
}

type Reat1 struct {
	leftUp, rightDown Point
}
type Reat2 struct {
	leftUp, rightDown *Point
}


type A struct {
	Num int
}

type  B struct {
	Num int
}

type  Student struct {
	Name string
	Age int
}

type Stu Student

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

// 细节说明
// 结构体和结构体变量(实例)的区别和联系
// 1.结构体是自定义的数据类型，代表一类事物
// 2.结构体变量是具体的，时机的，代表一个具体变量

// 如何声明结构体
// type 结构体名称 struct{}

// 从概念或者叫法上看：结构体字段 = 属性 = field
// 字段是结构体的一个组成部分，一般是基础数据类型、数组，也可能是引用类型
// 比如我们前面定义结构体的Name string就是属性

// 注意事项
// 1.字段声明语法同变量，示例：字段名 字段类型
// 2.字段的类型可以为:基本数据类型、数组或引用类型
// 3.在创建一个结构体变量后，如果没有给字段赋值，都对应一个默认值
// 指针，slice，map的默认值都是nil，即还没有分配空间
// 4.不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，不影响另外一个，结构体是值类型
















