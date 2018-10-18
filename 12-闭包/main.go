
package main 

import(
	"fmt"
	"strings"
)

func main() {
	// 闭包
	// 闭包就是一个函数和与其相关的引用环境组合的一个整体(实体)

	f1 := AddUpper()

	fmt.Println(f1(1))
	fmt.Println(f1(2))
	fmt.Println(f1(3))

	f2 := AddUpper2()

	fmt.Println(f2(1))
	fmt.Println(f2(2))
	fmt.Println(f2(3))

	f3 := makeSuffix(".jpg")
	fmt.Println("文件处理后",f3("winter"))
	fmt.Println("文件处理后",f3("spring.jpg"))

}
// 细节说明
// 1.AddUpper是一个函数，返回的数据类型是 func(int) int
// 2.返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和n形成一个整体，构成闭包
// 3.大家可以这样理解：闭包是类，函数式操作，n是字段，函数和他使用到n构成闭包
// 4.当我们反复的调用f1函数时，因为n是初始化一次，因此每次调用一次就进行累计
// 5.我们要搞清楚闭包的关键，就是要分析出返回的函数它使用（引用）到那些变量，因为函数和他引用到的变量共同构成闭包
func AddUpper() func (int) int {
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
	
}

func AddUpper2() func (int) int {
	var n int = 20
	var str string = "hello"

	return func (x int) int {
		n = n + x
		str += string(36)
		fmt.Println("str=",str)
		return n
	}
	
}


// 闭包案例
func makeSuffix(suffix string) func (string) string {
	// 如果name没有制定后缀，就加上，否则返回原来的名字
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}








