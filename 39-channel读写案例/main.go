

package main 

import (
	"fmt"
	"time"
)


// 案例
func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)

	go readData(intChan, exitChan)


	time.Sleep(time.Second * 5)

	for {
		_, ok := <- exitChan
		if !ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	fmt.Println("W")
	for i := 0; i < 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("w--",i)
	}

	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	fmt.Println("R")
	for {
		v, ok := <- intChan

		if !ok {
			break
		}
		fmt.Printf("r--读到数据=%v\n",v)
	}

	exitChan <- true
	close(exitChan)
}

// 如果只向管道写入数据，而没有读取，就会出现阻塞而deadlock








