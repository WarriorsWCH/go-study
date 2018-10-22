

package main 

import(
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 0; i <= 80000; i++ {
		intChan <- i
	}

	close(intChan)
}


// 从intChan中取出数据，判断是否为素数，是的话就放到primeChan
func prinmeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <- intChan
		// 取不到数据
		if !ok {
			break
		}

		flag = true
		// 判断是否是素数
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("一个协程退出")
	// 这里还不能关闭primeChan
	exitChan <- true
}
func main(){

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool,8)

	start := time.Now().Unix()

	go putNum(intChan)

	for i := 0; i < 8; i++ {
		go prinmeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<- exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=",end - start)
		// 8个协程都退出 关闭chan
		close(primeChan)
	}()

	for {
		v, ok := <- primeChan

		if !ok {
			break
		}
		fmt.Println(v)
	}

	fmt.Println("主线程退出")
}




















