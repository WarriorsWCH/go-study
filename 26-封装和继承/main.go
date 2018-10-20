

package main 

import (
	"fmt"
)


// 封装
// 封装就是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，
// 程序的其他包只有通过被授权的操作（方法），才能对字段进行操作
// go的体现：
// 1.对结构体中的属性进行封装
// 2.通过方法，包实现封装

// 继承
// 继承可以解决代码复用，让我们的编程更加靠近人类思维
// 多个结构体存在相同的属性（字段）和方法时
// 可以从这些结构体重抽象出结构体（比如刚才的Student），在该结构体中定义这些相同属性和方法
// 其他的结构体不需要重新定义这些属性（字段）和方法，只需嵌套一个Student匿名结构体即可

type Goods struct {
	Name string
	Price int
}

type Book struct {
	Goods //这里就是嵌套匿名结构体 Goods
}

func (g *Goods) GetName() string {
	return g.Name
}

func (b *Book) GetName() string {
	return "ok"
}

// 继承给编程带来的便利
// 代码的复用性提高了
// 代码的扩展性和维护性提高了

// 继承细节
// 结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法，都可以使用
func main() {
	var b Book
	b.Goods.Name = "数学"
	b.Goods.Price = 20

	name := b.Goods.GetName()
	fmt.Println(name)//数学

	// 匿名结构体字段访问可以简化
	b.Name = "英语"
	b.Price = 10
	name = b.GetName()
	fmt.Println(name)//ok
	name = b.Goods.GetName()
	fmt.Println(name)//英语

	// 1.当我们自己接通过b访问字段或方法时，执行流程如下，比如b.Name
	// 编译器会先看b对应的类型有没有Name，如果有，则直接调用B类型的Name字段
	// 2.如果没有就去看B中嵌入的匿名结构体A有没有声明Name字段，如果有就调用，如果都找不到就报错
	// 3.当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问
	// 如希望访问，匿名结构体的字段和方法，可以通过匿名结构体名来区分
	// 4.结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法
	// （同时结构体本身没有相同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错
	// 5.如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系
	// name在访问组合的结构体的字段或方法时，必须带上结构体的名字

	tv := TV{Goods{"电视机",500}, Brand{"海尔","山东"}}
	tv2 := TV{
		Goods{
			Name: "洗衣机",
			Price: 600,
		},
		Brand{
			Name:"小米",
			Address:"北京",
		},
	}
	fmt.Println(tv,tv2)

	tv3 := TV2{&Goods{"音响",300},&Brand{"腾讯","深圳"}}
	tv4 := TV2{
		&Goods{
			Name:"录音机",
			Price:200,
		},
		&Brand{
			Name:"阿里",
			Address:"杭州",
		},
	}
	fmt.Println(tv3,tv4)
	//{0xc42000a060 0xc42000a080} {0xc42000a0a0 0xc42000a0c0}
	fmt.Println(tv3.Goods,tv4.Goods)
	//&{音响 300} &{录音机 200}
	fmt.Println(*tv3.Goods,*tv4.Goods)
	//{音响 300} {录音机 200}

	// 结构体的匿名字段是基本数据类型，如何访问，下面代码输出什么
	var e E
	e.Name = "haha"
	e.Age = 33
	e.int = 1
	e.n = 2
	fmt.Println(e)
	//{{haha 33} 1 2}
	// 如果一个结构体有int类型的匿名字段，就不能有第二个
	// 如果需要有多个int字段，则必须给int字段指定名字
}

// 6.嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age int
}

type E struct {
	Monster
	int
	n int
}


// 细节说明：
// 封装的实现步骤
// 1.讲结构体、字段（属性）的首字母小写（不能导出了，其他包不能使用），类似private
// 2.给结构体所在包提供一个工厂模式的函数，首字母大写，类似构造函数
// 3.提供一个首字母大写的Set方法，类似public，用于对属性判断并赋值
// 4.提供一个首字母大写的Get方法，用于获取属性的值
// 在go开发中并没有特别强调封装，这点不像java
// 不用总是用java的语法特性来看待go，go本身对面向对象的特性做了简化的

// 如一个struct嵌套了多个匿名结构体，name该结构体可以直接访问嵌套的匿名结构体的字段和方法，从而实现了多重继承
// 如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来区分
// 为了保证代码的简洁性，建议大家尽量不使用多重继承


















