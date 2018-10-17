
package main

import(
	"fmt"
)

// 指针
// 1.基本数据类型，变量存的就是值，也叫值类型
// 2.获取变量的地址，用&


func main (){
	var i int = 10
	fmt.Println("i的地址=",&i)//0xc4200160b0

	// ptr是一个指针变量
	// ptr的类型是*int
	// ptr本身的值是&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n",ptr)//0xc4200160b0
	fmt.Printf("ptr的地址=%v\n",&ptr)//0xc42000c030
	fmt.Printf("ptr指向的值=%v\n",*ptr)//10

	ptrShow()
	ptrShow2()
}

func ptrShow(){
	var num int = 999
	var ptr *int = &num
	*ptr = 1000
	fmt.Println("num的地址%v\n",&num)
	fmt.Printf("num的值=%v\n",num)
}

func ptrShow2(){
	a := 100
	b := 200
	var ptr *int = &a
	*ptr = 300
	ptr = &b
	*ptr = 400
	fmt.Printf("a=%d b=%d ptr=%d\n",a,b,*ptr)
}

// 指针使用细节
// 1.值类型都有对应的指针类型，形式为*数据类型
// float32 对应的指针类型就是 *float，依次类推
// 2.值类型包括：基本数据类型 int..，float..，bool，string，数组和结构体



// 值类型和引用类型
// 值类型包括：基本数据类型 int..，float..，bool，string，数组和结构体
// 引用类型：指针，slice切片，map，管道chan，interface等都是引用类型

// 值类型：变量直接存储，内存通常在栈中分配
// 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值）
// 内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由GC回收


// 标识符
// 1.go对各种变量，方法，函数等命名时使用的字符序列成为标识符
// 2.凡是自己可以起名字的地方都叫标识符
// 命名规则
// 1.由26个英文字母大小写，0-9，_组成
// 2.数字不可以开头
// 3.go中严格区分大小写
// 4.标识符不能包含空格
// 5.下划线_本身在go中是一个特殊的标识符，成为空标识符
// 6.不能以系统保留关键字作为标识符

// 标识符注意事项
// 1.包名：保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和标准库冲突
// 2.变量名，函数名，常量名：采用驼峰法
// 3.如果变量名，函数名，常量名首字母大写，则可以被其他的包访问，如果首字母小写，则只能在本包中使用

//系统保留关键字
//break default func interface select case defer go map struct chan else goto package switch
//const fallthrough if range type continue for import return var

//系统预定义标识符
//append bool byte cap close complex complex64 complex128 uint16 copy false float32 float64 imag
//int8 int16 uint32 int32 int64 iota len make new nil panic uint64 print println real recover
//string true uint unit8 uintprt 











































