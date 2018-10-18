

package main 


import(
	"fmt"
)

var age = test()

func test() int {
	fmt.Println("test")
	return 100
}
// init函数
// 每一个源文件都可以包含一个init函数，该函数会在main函数执行前
// 被go运行框架调用，也就是说init会在main函数前被调用
func init() {
	fmt.Println("init")
	// 细节说明
	// 1.如果一个文件同时包含全局变量定义，init函数和main函数，则执行的流程是 全局变量定义 ->init函数 -> mian函数
	// 2.init函数最主要的作用，就是完成一些初始化的工作
	// 面试题：如果main.go和utils.go都含有变量定义，init函数时，执行的流程又是怎样的呢？
}

var (
	Func = func (n1 int, n2 int) int {
		return n1 * n2
	}
)
func main() {
	fmt.Println("age=",age)

	// 匿名函数
	// go支持匿名函数，匿名函数就是没有名字的函数，如果我们某个函数只是希望使用一次，
	// 可以考虑使用匿名函数，匿名函数也可以实现多次调用

	// 匿名函数的使用方式
	// 1.在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	r1 := func (n1 int, n2 int) int {
		return n1 + n2
	}(10,20)
	fmt.Println("r1=",r1)

	// 2.将匿名函数赋给一个变量（函数变量），再通过该变量来调用匿名函数
	myFunc := func (n1 int, n2 int) int {
		return n1 + n2
	}
	r2 := myFunc(10,30)
	fmt.Println("r2=",r2)

	// 全局匿名函数
	// 如果将匿名函数赋给一个全局变量，name这个匿名函数，就成为一个全局匿名函数，可以在程序有效
	r3 := Func(10,30)
	fmt.Println("r3=",r3)
}













