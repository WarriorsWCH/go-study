
// go两种执行方式的区别？
// 1.如果我们先编译生成了可执行文件，这个文件可以处处运行
// 2.go run 方式需要安装go的sdk
// 3.编译的可执行文件会大一些

// 编译指定输出名称为output
// go build -o output

// go编程的注意点
// 1.文件以go扩展名结尾
// 2.应用程序的执行入口是main()函数
// 3.go语言严格区分大小写
// 4.go每条语句没有分号
// 5.go是逐行编译的，一行只能写一条语句
// 6.定义的变量或者import没有用到，go编译不会通过
// 7.大括号成对出现，缺一不可

package main

import (
	"fmt"
	"unsafe"
)
// go变量使用的三种方式


// 定义全局变量
var n1 = 100
var name1 = "groot"

// 也可以一次性声明
var (
	n2 = 200
	name2 = "w3c"
)

func main() {
	fmt.Println("hello world")
	
	// go怨言的转义字符
	// \t一个制表位
	// \n换行
	// \\一个\
	// \"一个"
	// \r表示从当前行最前面开始覆盖掉之前的内容
	fmt.Println("hello\tworld")//hello	world
	fmt.Println("hello\nworld")
	fmt.Println("C:\\usr\\go")//C:\usr\go
	fmt.Println("go say \"wwo\"")//go say "wwo"
	fmt.Println("go say \rwwo")//wwosay


	// 1.指定变量类型，声明后如不赋值，使用默认值，如int的默认值是0
	var i int
	fmt.Println("i=",i)
	// 2.根据值自动判断类型（类型推导）
	var num = 10.11
	fmt.Println("num=",num)
	// 3.省略var 注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
	name := "mark"
	fmt.Println("name=",name)
	// 4.一次声明多个变量
	var n1,n2,n3 int
	fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)
	// 5.第二种一次性声明多个变量
	var n4,name2,n5 = 100, "tom", 888
	fmt.Println("n4=",n4,"name2=",name2,"n5=",n5)
	// 6.一次性声明多个变量同样可以使用类型推导
	n6, name3 := 200, "jack"
	fmt.Println("n6=",n6,"name3=",name3)
	// 7.该区域的数据可以在范围内不断变化（不能改变数据类型）
	n4 = 200
	n4 = 300
	fmt.Println("n4=",n4)
	// 8.变量在一个函数或者代码块种不能重名
	// 9.变量= 变量 + 值 + 数据类型
	// 10.变量如果没有赋予初始值，int默认值为0，sting默认值为空串，小数默认为0
	// 11.关于+当变量都是数值类型的时候做加法操作，当都是字符串的时候作字符串链接操作


	// 数据类型
	// --基本数据类型：
	// 1.数值型（int，int8，int16，int32，int64，uint，uint4，uint8，uint16，uint32，uint64，byte，float32，float64）
	// 2.字符型，没有专门的字符型，使用byte来保存单个字母字符
	// 3.布尔型bool
	// 4.字符串string
	// --复杂数据类型：
	// 指针，数组，结构体，管道，函数，切片，接口，map

	var i1 int = 1
	fmt.Println("i1=",i1)
	var j1 int8 = 127
	fmt.Println("j1=",j1)

	// int 有符号，32位系统占用4字节，可以表示 负2的31次方 到 2的31次方-1，64位占用8字节，可以表示负2的63次方到2的63次方-1
	// uint 无符号，32位系统占用4字节，可以表示 0到2的32次方-1.......
	// rune 有符号，位用4字节，可以表示 负2的31次方到 2的31次方-1，等价int32表示一个unicode编码
	// byte，无符号，与uint8等价，表示0-255，当腰存储字符的时候用byte

	var a int = 233
	fmt.Println("a=",a)
	var b uint = 233233
	fmt.Println("b=",b)
	var c byte = 255
	fmt.Println("c=",c)

	// 使用细节
	// 1.go各整数类型分有符号无符号，int uint的小大小和系统无关
	// 2.go的整数默认声明为int型

	var a1 = 100
	fmt.Printf("a1的类型%T \n",a1)
	//查看某个变量占用的字节大小和数据类型
	var a2 int64 = 10
	fmt.Printf("a2的类型是%T a2占用的字节数是%d",a2, unsafe.Sizeof(a2))//字节是8
	// go程序中整数变量在使用时，遵循保小不保大的原则，即在保证程序运行正常的情况下，尽量使用占用空间小的数据类型

}










