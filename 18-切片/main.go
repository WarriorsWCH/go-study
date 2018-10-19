

package main 


import (
	"fmt"
)

func main() {
	// 切片
	// 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递机制
	// 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度都一样
	// 切片的长度是可以变化的，因此切片是一个可以动态变化的数组

	var intArr [5]int = [...]int{1,2,3,44,55}
	slice := intArr[1:3]
	fmt.Println("intArr=",intArr)
	fmt.Println("slice的元素是：",slice)//[2 3]
	fmt.Println("slice的个数是：",len(slice))//2
	fmt.Println("slice的容量是：",cap(slice))//4
	// intArr[1:3]表示slice引用到intArr这个数组
	// 引用intArr数组的起始下标是1，结束是3但是不包含3
	fmt.Printf("intArr%p slice%p\n",&intArr,&slice)
	// intArr0xc42001a180 slice0xc42000a060


	// 第二种方式
	// 通过make来创建切片
	//基本语法:var 切片名 []type = make([]type, len, [cap])
	//参数说明: type: 就是数据类型 len : 大小 cap :指定切片容量，可选，如果你分配了cap,则要求 cap>=len.
	var slice2 []float64 = make([]float64,5,10)
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println("slice2",slice2)//[0 10 0 20 0]
	// 说明
	// 1.通过make方式创建切片可以指定切片的大小和容量
	// 2.如果没有给切片的各个元素赋值，那么就会使用默认值
	// 3.通过make方式创建的切片对应的数组是由make底层委会，对外不可见，即只能通过slice去访问各个元素


	// 方式三
	// 定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var slice3 []string = []string{"tom","jack","tony"}
	fmt.Println("slice3=",slice3,"len=",len(slice3),"cap=",cap(slice3))

	// 方式一和方式二的区别
	// 方式一是直接饮用数组，这个数组是事先就存在的，程序是可见的
	// 方式二是通过make来创建切片，make也会创建一个数组，是由切盼在底层维护的，程序员是看不见的



	// 切片的使用细节
	// 1，切片初始化时,仍然不能越界，范围是[0-len(arr)]之间，但是可以动态增长
	// 2.cap是一个内置函数，用于统计的容量，即最大可以存放多少个元素
	// 3.切片定义完后，还不能使用，因为本身是一个空，需要让其引用到一个数组，或者make一个空间供切片来使用
	// 4.切片可以继续切片

	var arr [5]int = [...]int{10,20,30,40,50}
	slice4 := arr[1:4]
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("slice4[%v]=%v\n",i,slice4[i])
	}

	slice5 := slice4[1:2]
	slice5[0] = 100
	fmt.Println("slice4",slice4,"slice5",slice5,"arr",arr)
	// 因为arr slice4 scli5指向的数据空间是同一个，所以改一个值其他都变了



	// 6.用append内置函数，可以对切片进行动态追加

	var slice6 [] int = []int{100,200,300}
	fmt.Printf("slice6=%p\n",&slice6)//0xc42000a120
	slice6 = append(slice6,400,500,600)
	fmt.Println("slice6",slice6)
	fmt.Printf("slice6=%p\n",&slice6)//0xc42000a120

	slice6 = append(slice6,slice6...)
	fmt.Println("slice6=",slice6)
	fmt.Printf("slice6=%p\n",&slice6)//0xc42000a120
	// 切片append操作的底层原理分析
	// 切片append操作的本质就是对数组扩容
	// go底层会创建一个新的数组，讲slice原来包含的元素拷贝到新的数组,slice重新引用到新数组


	// 切片的拷贝操作
	var slice7 []int = []int{1,2,3,4,5,6,7,8}
	var slice8 = make([]int,10)
	copy(slice8,slice7)
	fmt.Println("slice7:",slice7,"slice8:",slice8)
	fmt.Printf("slice7=%p  slice8=%p\n",&slice7,&slice8)


	var a []int = []int{1,2,3,4}
	var slice9 = make([]int,1)
	fmt.Println(slice9)
	copy(slice9,a)
	fmt.Println(slice9)

	// 1.slice是一个引用类型
	// 2.slice从底层来说，其实就是一个数据结构(struct结构体)
	//type slice struct {
	//ptr *[2]int 
	//len int 
	//cap int
	//}


}



























