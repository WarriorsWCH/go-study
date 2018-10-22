

package main 

import(
	"fmt"
	"time"
)

// goroutine中使用recover，解决协程中出现panic，导致程序崩溃的问题
func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}

}

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=",i)
		time.Sleep(time.Second)
	}
}

func test() {
	// 这里我们可以使用 defer + recover
	defer func() {
		// 捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	// 定义了一个map
	var myMap map[int]string
	myMap[0] = "gogogo"// panic: assignment to entry in nil map
	// myMap实际上定义了一个map指针，这个指针指向NULL。
	// 既然myMap是一个空指针，并没有一个真实的map存在，所以也就不能对myMap进行内存访问操作
}
