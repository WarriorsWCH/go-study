

package main

import (
	"fmt"
	"strconv"
)

// 基本数据类型的相互转换
// go和java/C不同，go在不同类型的变量之间赋值时需要显示转换，也就是说go中数据类型不能自动转换
// 如：int(10.0)

// 注意事项
// 1.go中，数据类型转换可以是表示范围小--->表示范围大，也可以范围大--->范围小
// 2.被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化
// 3.在转换中，比如将int64转成int8，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样，因此在转换时，需要考虑范围

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n",i,n1,n2,n3)
	fmt.Printf("i type is %T\n",i)
	fmt.Printf("n3 type is %T\n",n3)

	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=",num2)

	// 基本数据类型和string的转换
	// 1.基本类型转string类型 fmt.Sprintf("%参数",表达式)
	str := fmt.Sprintf("%d",i)
	fmt.Printf("str type %T str %q\n",str,str)

	str2 := fmt.Sprintf("%f",n1)
	fmt.Printf("str2 type %T str2 %q\n",str2,str2)

	isB := false
	str3 := fmt.Sprintf("%t",isB)
	fmt.Printf("str3 type %T str3 %q\n",str3,str3)

	c := 'c'
	str4 := fmt.Sprintf("%c",c)
	fmt.Printf("str4 type %T str4 %q\n",str4,str4)


	// 使用strconv包的函数
	num3 := 99
	num4 := 23.22
	isL := true

	str5 := strconv.FormatInt(int64(num3),10)
	fmt.Printf("str5 type %T str5 %q\n",str5,str5)

	// f格式，10表示小数保留10位，64表示这小数是float64
	str6 := strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str6 type %T str6 %q\n",str6,str6)

	str7 := strconv.FormatBool(isL)
	fmt.Printf("str7 type %T str7 %q\n",str7,str7)

	var num5 int64 = 8765
	str8 := strconv.Itoa(int(num5))
	fmt.Printf("sr8 type %T str8 %q\n",str8,str8)



	// 2.string类型转基本类型
	stringToBase()

}

func stringToBase() {
	var str string = "true"
	var b bool

	// ParseBool会返回两个值（value bool, err error
	// 我们只想获取value，不像获取err所以_忽略
	b,_ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n",b,b)

	var str2 string = "12345678"
	var n1 int64
	var n2 int

	n1,_ = strconv.ParseInt(str2,10,64)
	n2 = int(n1)
	fmt.Printf("n1 type %T n1=%v\n",n1,n1)
	fmt.Printf("n2 type %T n2=%v\n",n2,n2)

	var str3 string = "12345.12345"
	var f float64
	f,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("f type %T f=%v\n",f,f)

	// string转基本数据类型的注意事项
	// 在将string类型转成基本数据类型时，要确保string类型能够转成有效的数据
	// 比如，我们可以把"123"转成一个整数，但是不能把"hello"转成一个整数，如果这样做go直接将其转成0
	// 其他类型也是一样的道理，float->0 bool ->false
}

















