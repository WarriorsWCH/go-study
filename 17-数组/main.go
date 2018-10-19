

package main 


import (
	"fmt"
	"time"
	"math/rand"
)

func main(){
	// 数组
	// 在go中，数组是值类型
	var hens [7]float64
	hens[0] = 0.0
	hens[1] = 1.0
	hens[2] = 2.0
	hens[3] = 3.0
	hens[4] = 4.0
	hens[5] = 5.0
	hens[6] = 6.0

	totalWight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWight += hens[i]
	}
	fmt.Println(totalWight)

	var arr [3]int
	fmt.Println(arr)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Printf("arr的地址是%p arr 0 1 2的地址是%p %p %p\n",&arr,&arr[0],&arr[1],&arr[2])


	// 四种初始化数组的方式
	var arr1 [3]int = [3]int{1,2,3}
	fmt.Println("arr1=",arr1)

	var arr2 = [3]int{4,5,6}
	fmt.Println("arr2=",arr2)

	var arr3 = [...]int{7,8,9}
	fmt.Println("arr3=",arr3)

	var arr4 = [...]int{1:800,0:900,2:999}
	fmt.Println("arr4=",arr4)


	// 数组遍历
	// for-range遍历
	// 1.第一个返回值index是下标
	// 2.第二个返回值value是該下标位置的值
	// 3.他们都是for循环内部可见的变量
	// 4.如果不想使用index可以使用_代替
	// 5.index和value可以取其他名字

	for _, val := range arr4 {
		fmt.Println("值为：",val)
	}

	// 数组注意事项
	// 1.数组是多个相同类型数据的组合，一个数组一旦声明/定义了，其长度是固定的，不能动态变化
	// 2.vae arr []int 这时 arr就是一个slice切片
	// 3.数组中的元素可以使任何数据类型，包括值类型和引用类型，但是不能混用
	// 4.数组创建后，如果没有赋值，有默认值
	// 5.使用步骤 1.声明数组并开辟空间 2.给数组哥哥元素赋值 3.使用数组
	// 6.数组的下标从0开始
	// 7.数组下标必须在指定范围内使用，否则报panic:数组越界
	// 8.go的数组属值类型，在默认情况下是值传递，因此会进行拷贝，数组间不会相互影响
	// 9.长度是数组类型的一部分，在传递参数时，需要考虑数组的长度

	var myArr [3]int = [3]int{1,2,3}
	test1(myArr)
	fmt.Println("myarr=",myArr)
	test2(&myArr)
	fmt.Println("myarr=",myArr)


	func1()
	func2()
	func3()
}

func test1(arr [3]int) {
	arr[0] = 9
	fmt.Println("test1-arr=",arr)
}

func  test2(arr *[3]int) {
	arr[0] = 9

	fmt.Println("test2-arr=",arr);
	fmt.Printf("%p %v %p\n",&arr,*arr,&arr[1])
}

// 细节说明
// 数组的地址可以通过数组名来获取 &arr
// 数组的第一个元素的地址，就是数组的首地址
// 数组的各个元素的地址间隔依据数组的类型决定带的，比如 int64 --> 8 int32 --> 4.....


// 数组的应用
// 创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印出来
// 提示:字符数据运算 'A'+1 -> 'B'
func func1() {
	var ele [26]byte
	for ind, _ := range ele{
		ele[ind] = 'A' + byte(ind)
	}
	for _, val := range ele{
		fmt.Printf("%c,",val)
	}

	fmt.Println()
}



// 请求出一个数组的最大值，并得到对应的下标。

func func2() {
	var intArr [6]int = [...]int{12,4,23,56,87,56}
	maxVal := intArr[0]
	maxIndex := 0
	for ind, val := range intArr {
		if val > maxVal {
			maxVal = val
			maxIndex = ind
		}
	}
	fmt.Printf("maxVal=%d,maxIndex=%d\n",maxVal,maxIndex)
}


// //随机生成五个数，并将其反转打印
func func3() {
	var intArr3 [5]int
	len := len(intArr3)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}

	fmt.Println("交换前,",intArr3)
	for i := 0; i < len/2; i++ {
		intArr3[len - 1 - i], intArr3[i] = intArr3[i], intArr3[len - 1 - i]
	}
	fmt.Println("交换后，",intArr3)
}


















