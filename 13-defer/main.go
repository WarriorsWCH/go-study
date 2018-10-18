
package main 

import (
	"fmt"
)

func main() {
	
	// defer
	// 在函数中，程序员经常需要创建资源(比如：数据库连接、文件句柄、锁等)，为了在函数执行完毕后，及时的释放资源
	res := sum(10,20)
	fmt.Println("haha")
	fmt.Println("res=",res)
}

func sum(n1 int, n2 int) int {
	fmt.Println("sum")
	// 当执行到defer暂时不执行，会将defer后的语句压如defer栈
	// 当函数执行完毕后，按照先入后出的方式依次执行
	defer fmt.Println("ok n1",n1)
	defer fmt.Println("ok n2",n2)

	// defer值拷贝
	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok res",res)
	return res
}
// 细节说明
// 1.当go执行到一个defer时，不会立即执行defer后的语句，而是将defer后的语句压入到一个栈中，然后继续执行函数下一个语句
// 2.当函数执行完毕后，在从defer栈中，依次从栈顶去除语句执行(先入后出的机制)
// 3.在defer将语句放入到栈时，也会将相关的值拷贝同时入栈

// 最佳实战
// 1.defer最主要的价值在于，当函数执行完毕后，可以及时的释放函数创建的资源
// 2.在go编程中的通常做法是，创建资源后，比如打开了文件，可以执行defer file.close()
// 3.在defer将语句放入到栈时，也会将相关的值拷贝同时入栈后，可以继续使用创建资源
// 4.当函数完毕后，系统会依次从defer栈中，取出语句，关闭资源
// 5.这种机制，非常简洁，程序员不用再为在什么时候关闭资源而烦心