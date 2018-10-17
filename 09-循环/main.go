
package main 

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	// for循环
	var count int 
	for count = 0; count < 10; count++ {
		fmt.Println(count)
	}

	// 第二种写法
	// 也可以使用for ; ; {} 是哟个无限循环，通常配合break语句使用
	for count < 20 {
		fmt.Println("count=",count)
		count++
	}

	// go提供for-range的方式，可以方便遍历字符串和数组
	// 传统方式
	var str string = "hello world 呵呵" 
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c \n", str[i])
	}

	// for-range
	for index, val := range str {
		fmt.Printf("index=%d value=%c\n",index,val)
	}

	// 注意
	// 如果我们的字符串含有中文，name传统的遍历字符串方式是错误的，会出现乱码
	// 原因是传统的对字符串的遍历是按照字节来遍历，而一个汉字在utf8编码是对应3个字节
	// for-range是按照字符遍历的，所以没有问题
	
	// 如何解决：需要将str转成 []rune切片
	str2 := []rune(str)
	for i := 0; i <len(str2); i++ {
		fmt.Printf("%c \n", str2[i])
	}


	// go没有while和do while
	// 可以使用for取代

	// 九九乘法表
	var num byte
	for num = 1; num < 10; num++{
		var num2 byte
		for num2 = 1; num2 <= num; num2++ {
			fmt.Printf("%d * %d = %d\t",num2,num2,num*num2)
		}
		fmt.Println()
	}

	myGoTo()

}

func myGoTo() {
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1 //生成【0，100)
		fmt.Println("n=",n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99用了%d次\n",count)
	// break细节
	//break语句出现在多层嵌套语句块中时，可以通过标签指明要终止的是那一层语句块
	// break默认会跳出最近的for循环

	//continue
	//continue 语句用于结束本次循环，继续执行下一次循环。
	//continue 语句出现在多层嵌套的循环语句体中时，
	//可以通过标签指明要跳过的是哪一层循环, 这 个和前面的 break 标签的使用的规则一样.



	// goto
	// go语言的goto语句可以无条件的转移到程序中指定的行
	// goto语句通常与条件语句配合使用，可用来实现条件的转移，跳出循环体等功能
	// 在go程序设计中一般不主张使用goto语句，以免造成程序流程的混乱，是理解和调试程序都产生困难

	var n int = 30
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok2")
	label1:
	fmt.Println("ok3")
	// 如果return是在普通的函数，则表示跳出该函数，即不再执行函数中return后面的代码，也可以理解终止函数
	// 如果return是在main函数，表示终止main函数，也就是说终止程序
}






























