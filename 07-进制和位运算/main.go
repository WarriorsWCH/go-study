

package main

import(
	"fmt"
)

func main() {

	// output()


	// 进制
	// 对于整数，四种表示方式
	// 二进制：0，1，满2进1，在go中，不能直接使用二进制来表示一个整数，它沿用了c的特点
	// 十进制：0-9，满10进1
	// 八进制：0-7，满8进1，以数字0开头表示
	// 十六进制：0-9及A-F，满16进1，以0x或者0X开头表示，此处的A-F不区分大小写

	var i int = 5
	fmt.Printf("%b \n", i) //101

	var j int = 011
	fmt.Println("j=",j) // 9

	var k int = 0x11
	fmt.Println("k=",k) //17

	//二进制转八进制
	//将二进制数的每3位一组(从低到高)转成对应的八进制数即可
	//二进制转十六进制
	//将二进制的每4位一组转成对应的十六进制即可
	//八进制转二进制
	//将八进制的每一位转换成一个对应的三位二进制数即可
	//十六进制转二进制
	//将十六进制的每一位转换成对应的四位二进制数即可

	//对于有符号的数而言
	//1. 二进制的最高位是符号位：0表示正数 1表示负数
	//2. 正数的反码 原码 补码 都一样
	//3. 负数的反码=它的原码符号位不变，其他位取反
	// 1 => 00000001 00000001 00000001(补)
	//-1 => 10000001 11111110 11111111(补)
	//4. 负数的补码=他的反码+1
	//5. 0的反码补码都是0
	//6. 在计算机运算的时候，都是以补码的方式来运算的
	//计算机世界是没有减法的都是加法 1 - 1 =》 1 + (-1)
	

	//位运算演示
	fmt.Println("2&3=",2&3)
	fmt.Println("2|3=",2|3)
	fmt.Println("-2^2=",-2^2)

	//只要是运算都是补码来计算的
	//>>、<< 右移和左移,运算规则:
	//右移运算符 >>:低位溢出,符号位不变,并用符号位补溢出的高位
	//左移运算符 <<: 符号位不变,低位补 0

	fmt.Println("1>>2",1>>2)
	fmt.Println("1<<2=",1<<2)
}
func output() {
	// 键盘输入语句
	var (
		name string
		age byte
		salary float64
		isPass bool
	)

	fmt.Println("请舒服姓名 年龄 薪水 是否通过考试 请按空格隔开")
	fmt.Scanf("%s %d %f %t",&name,&age,&salary,&isPass)
	fmt.Printf("姓名：%v\t 年龄：%v\t 薪水：%v\t 是否通过：%v\t",name,age,salary,isPass)
}