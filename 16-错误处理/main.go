

package main 


import (
	"fmt"
	"errors"
)

func main() {
	// 错误处理
	// go中引入的处理方式为：defer，panic，recover
	// 这几个异常的使用场景可以这么简单描述：go中可以抛出一个panic的异常，
	// 然后在defer中，通过recover捕获这个异常，然后正常处理

	test()
	// test3()
	fmt.Println("继续执行")
}

func test() {
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10;
	num := 0
	res := num1/num
	fmt.Println("res=",res)
}

// 自定义错误的介绍
// go程序中，也支持自定义错误，使用errors.New 和 panic内置函数
// errors.New("错误说明"),会返回一个error类型的值，表示一个错误
// panic内置函数，接受一个interface{}类型的值（也就是任何值了）作为参数，可以接受error类型的变量，输出错误信息，并退出程序


func test2() error{
	return errors.New("发生错误")
}

func test3() {
	err := test2()
	if err !=nil {
		panic(err)
	}
}








