


package main 

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	// 字符串函数
	// 统计字符串的长度，按字节
	str := "hello你"
	fmt.Println("str len=",len(str))

	// 避免乱码
	str2 := "hello世界"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符：%c\n",r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("12")
	if err != nil {
		fmt.Println("转换错误",err)
	} else {
		fmt.Println("数字：%d\n",n)
	}
	// 整数转字符串
	str3 := strconv.Itoa(123456)
	fmt.Println(str3)

	// 字符串转[]byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n",bytes)

	// byte转字符串
	str3 = string([]byte{97,98,99})
	fmt.Printf("str3=%v,str3=%T\n",str3,str3)


	// 10进制转2，8，6进制
	num := strconv.FormatInt(123,2)
	fmt.Printf("num=%v,num=%T\n",num,num)

	// 查找子串是否在指定的字符串中
	isHas := strings.Contains("hello","ll")
	fmt.Printf("isHas=%v,isHas=%T\n",isHas,isHas)


	// 统计一个字符串有几个指定的子串
	n1 := strings.Count("gsahadgwuify","a")
	fmt.Printf("n1=%v,n1=%T\n",n1,n1)

	// 不区分大小写的字符串比较(==是区分大小写的)
	fmt.Println(strings.EqualFold("abc","Abc"))

	// 返回子串在字符串第一次出现的index值，如果没有返回-1
	fmt.Println(strings.Index("adeda","a"))


	// 返回子串在字符串最后一次出现的index，如果没有返回-1
	fmt.Println(strings.LastIndex("kshfiuefwfewj","f"))

	// 将指定的子串替换成另外一个子串
	// 可以指定希望替换几个，-1表示替换全部
	str4 := "go go run build"
	str5 := strings.Replace(str4, "go","c++",-1)
	fmt.Println(str4,str5)

	// 按照指定的某个字符，为分割标识，将一个字符串拆分成字符数组
	arr := strings.Split("a&b&c&d","&")
	fmt.Println(arr)

	// 将字符串的字母进项大小写转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("me"))

	// 去掉字符串两边的空格
	fmt.Println(strings.TrimSpace(" i am yours "))

	// 将字符串左右两边指定的字符去掉
	fmt.Println(strings.Trim("^哈哈^","^"))
	fmt.Println(strings.TrimLeft("^哈哈^","^"))
	fmt.Println(strings.TrimRight("^哈哈^","^"))

	// 判断字符串是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("http//:baidu.com","http"))

	// 判断字符串是否以指定字符结束
	fmt.Println(strings.HasSuffix("http//:baidu.com","com"))
}















