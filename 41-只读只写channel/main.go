

package main 

import (
	"fmt"
	"time"
)

func main() {

	// 1.channel可以声明为只读，或者只写性质

	// 只可以写
	var chan1 chan <- int
	chan1 = make(chan int, 3)
	chan1 <- 20
	fmt.Println("chan1=",chan1)

	// 只可以读
	// var chan2 <- chan int
	// num1 := <-chan2
	// fmt.Println("num1=",num1)

	// 默认双向
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0;
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束")

	// 使用select可以解决从管道取数据的阻塞问题
	// 1.定义一个管道10个数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定义一个管道5个数据
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++{
		stringChan <- "hello"+fmt.Sprintf("%d",i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致deadlock
	// 可以使用select方式解决
	for {
		select {
			case v := <-intChan:
				fmt.Printf("从intChan读取的数据%d\n",v)
				time.Sleep(time.Second)
			case v := <-stringChan:
				fmt.Printf("从stringChan读取的数据%s\n",v)
				time.Sleep(time.Second)
			default:
				return
		}
	}
	// select是go中的一个控制结构，类似switch语句，用于处理异步IO操作，select会监听case语句
	// 中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读能写）时，将会触发相应的动作
	// 1.如果有多个case都可以运行，select会随机公平的选出一个执行，其他不会执行
	// 2.如果没有可运行的case语句，且有default语句，那么就会执行default的动作
	// 3.如果没有可运行的case语句，且没有default语句，select将阻塞，知道某个case通信可以运行
}


func send(ch chan <- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	var a struct{}
	exitChan <- a
}

func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch

		if !ok {
			break
		} 
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
















