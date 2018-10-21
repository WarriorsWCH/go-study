

package main 

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	// flag包解析命令行参数
	// 基础方式
	fmt.Println("命令行参数有：",len(os.Args))

	for i,v := range os.Args {
		fmt.Println(i,v)
	}

	var (
		user string
		pwd string
		host string
		port string
	)

	// &user就是接受用户命令行中输入的-u后面的参数值
	// ”“默认参数
	// u 就是-u指定参数
	// "用户名,默认为空" 说明
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "用户密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机，默认为localhost")
	flag.StringVar(&port, "port", "3306", "端口，默认为3306")
	// 必须调用下
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v port=%v",user,pwd,host,port)
}

//测试
//Go 语言中自带有一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试
//testing 框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基
//于该框架写相应的压力测试用例。通过单元测试，可以解决如下问题:
//1.确保每个函数是可运行，并且运行结果是正确的
//2.确保写出来的代码性能是好的，
//3.单元测试能及时的发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决，
//而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保持稳定

//细节
//1.测试用例文件名必须以 _test.go 结尾。 比如 cal_test.go , cal 不是固定的。
//2.测试用例函数必须以 Test 开头，一般来说就是 Test+被测试的函数名，比如 TestAddUpper
//3.TestAddUpper(t *tesing.T) 的形参类型必须是 *testing.T
//4. 一个测试用例文件中，可以有多个测试用例函数，比如 TestAddUpper、TestSub
//5.运行测试用例指令
// go test [如果运行正确，无日志，错误时，会输出日志]
// go test -v [运行正确或是错误，都输出日志]
//6.当出现错误时，可以使用 t.Fatalf 来格式化输出错误信息，并退出程序
//7.t.Logf 方法可以输出相应的日志
//8.测试用例函数，并没有放在 main 函数中，也执行了，这就是测试用例的方便之处.
//9.PASS 表示测试用例运行成功，FAIL 表示测试用例运行失败
//10.测试单个文件，一定要带上被测试的原文件
//11.go test -v cal_test.go cal.go
//12.测试单个方法 go test -v -test.run TestAddUpper


















