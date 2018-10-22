

package main 

import (
	"fmt"
)


// channel本质就是一个数据结构-队列
// 数据时先进先出
// 线程安全，多goroutline访问时，不需要加锁，就是说channel本身就是线程安全的
// channel有类型，一个string的channel只能存放sting类型数据

type Person struct{}

var (
	intChan chan int
	mapChan chan map[int]string
	personChan chan Person 
	personChan2 chan *Person 
)

type Cat struct {
	Name string
	Age int 
}
func main() {
	// 创建一个可以存放3个int类型的管道
	intChan = make(chan int, 3)
	fmt.Printf("intChan的值=%v intChan本身的地址=%p\n",intChan,&intChan)

	// 写数据
	intChan <- 20
	num := 211

	intChan <- num
	intChan <- 50
	// 看看管道的长度和cap容量
	fmt.Printf("channel len=%v cap=%v \n",len(intChan),cap(intChan))
	// fmt.Println(*intChan)//invalid indirect of intChan (type chan int)

	// 读数据
	var num2 int 
	num2 = <-intChan
	fmt.Println("num2=",num2)
	fmt.Printf("channel len=%v cap=%v \n",len(intChan),cap(intChan))

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告deadlock
	num3 := <-intChan
	num4 := <-intChan

	fmt.Println(num3,num4)

	// num5 := <- intChan
	// fmt.Println(num5)
	// all goroutines are asleep - deadlock!
	
	// channel中只能存放指定的数据类型
	// channel的数据存放满后，就不能再存放了
	// 如果从channel取出数据后，可以继续放入
	// 在没有使用协程的情况下，如果channel数据去玩了，再取，就会报错dead lock

	allChan := make(chan interface{}, 3)
	allChan <- 3
	allChan <- "jack"
	cat := Cat{"小小",20,}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%T newCat=%v\n", newCat,newCat)
	a := newCat.(Cat)//使用类型断言
	fmt.Printf("newCat.Name=%v\n",a.Name)

	// channel的遍历与关闭
	// 使用内置函数close可以关闭channel，当channel关闭后，就不能再向channel写数据了，但是仍然可以从该channel读取数据

	close(allChan)
	// allChan <- 20//send on closed channel

	allChan2 := make(chan interface{}, 3)
	allChan2 <- 33
	allChan2 <- "TOM"
	cat = Cat{"大大",88,}
	allChan2 <- cat

	close(allChan2)

	for v := range allChan2 {
		fmt.Println("channel->",v)
	}
	// channel的遍历
	// channel支持for-range的方式遍历
	// 1.在遍历时，如果channel没有关闭，则会出现dead lock的错误
	// 2.咋遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会推出遍历
	// 3.for-range没有index，只有value

	// 细节说明
	// 1.channel是引用类型
	// 2.channel必须初始化才能写数据，即nake后才能使用
	// 3.管道是有类型的，intChan只能写如入证书类型int

}
















