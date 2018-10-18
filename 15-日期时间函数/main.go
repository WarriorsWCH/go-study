

package main 

import (
	"fmt"
	"time"
)

func main(){
	// 时间和日期函数

	// 当前时间
	now := time.Now()
	fmt.Printf("now=%v type=%T\n",now,now)//ype=time.Time

	// 获取年月日时分秒
	fmt.Printf("年=%v\n",now.Year())
	fmt.Printf("月=%v -- %v\n",now.Month(),int(now.Month()))
	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	// 格式化日期
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	// 这个字符串的各个数字是固定的，必须要这样写
	// 这个字符串的各个数字可以自由组合，这样可以按照程序需求来返回时间
	fmt.Printf(now.Format("15|04|05"))

	// 时间常量
	i := 1
	for {
		i++
		fmt.Println(i)
		// 休眠
		time.Sleep(time.Microsecond * 50000)

		if i == 100 {
			break
		}
	}

	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n",now.Unix(),now.UnixNano())
	// Unix时间，从1970年1月1日
	// unixnano纳秒数，如果纳秒位单位时间超过了int64表示范围，结果就会是未定义的

	start := time.Now().Unix()
	end := time.Now().Unix()
	fmt.Printf("耗费时间%d s\n",end - start)


	num1 := 100
	fmt.Printf("num1的类型%T,num1的值%v,num1的地址%v\n",num1,num1,&num1)

	num2 := new(int)
	*num2 = 200
	fmt.Printf("num2的类型%T,num2的值%v,num2的地址%v,num2这个指针指向的值%v\n",num2,num2,&num2,*num2)

	// 内置函数
	// go设计者为了编程方便，提供了一些函数，这些函数可以直接使用，我们称之为内置函数
	// 文档:https://studygolang.com/pkgdoc -> builtin
	// 1.len:用来求长度，比如string、array、slice、map、channel
	// 2.new:用来分配内存，主要用来分配值类型，比如int，float64...返回的是指针
	// 3.make：用来分配内存，主要是分配引用类型，比如channel、map....
}





















