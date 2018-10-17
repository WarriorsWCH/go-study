

package main

import (
	"fmt"
)

func main() {
	// string类型
	// go的字符串是由单个字节连接起来的
	// go语言的字符串的字节使用UTF-8编码标识unicode文本

	var address string = "hello"
	fmt.Println(address)
	fmt.Println(address[0])//104 打印的是h的utf-8编码值
	// 注意事项
	// 1.go语言的字符串的字节使用UTF-8编码标识unicode文本
	// 这样go统一使用UTF-8编码，中文乱码问题不会再困扰程序员了
	// 2.字符串一旦赋值了，字符串就不能修改了，go的字符串是不可变的
	// address[0] = 's'会报错
	// 3.字符串的两种表示形式
	// 3.1双引号，会识别转义字符
	// 3.2反引号，义字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果

	str := `
		func main() {
			
		}
	`
	fmt.Println(str)

	// 4.字符串拼接
	str2 := "hello"
	str2 += "world"
	fmt.Println(str2)

	// 5.当一行字符串太长时，需要使用多行字符串，可以如下处理
	str3 := "qwertyuiop"+
			"asdfghjkl"+
			"zxcvbnm"
	fmt.Println(str3)


	// 基本类型的默认值
	// 整数 0
	// 浮点型 0
	// 字符型 ""
	// 布尔类型 false
	var a int
	var b float32
	var c string
	var isb bool
	fmt.Println("a=",a,"b=",b,"c=",c,"isb=",isb)
}










