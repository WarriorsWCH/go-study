

package main 

import(
	"fmt"
)

func  main() {
	// 二维数组
	// 使用方式1
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j :=0; j < 6; j++{
			fmt.Print(arr[i][j]," ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n",&arr2[0])//0xc42001a180
	fmt.Printf("arr2[1]的地址%p\n",&arr2[1])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])//0xc42001a180
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])


	// 使用方式二
	// 直接初始化
	var arr3 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr3)
	// 二维数组在声明/定义时也对应有四种写法
	//var 数组名 [大小][大小]类型 = [大小][大小]类型{{初值..},{初值..}}
	//var 数组名 [大小][大小]类型 = [...][大小]类型{{初值..},{初值..}}
	//var 数组名 = [大小][大小]类型{{初值..},{初值..}}
	//var 数组名 = [...][大小]类型{{初值..},{初值..}}
}