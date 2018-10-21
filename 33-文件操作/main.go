
package main 

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

// 文件操作
func main(){
	file, err := os.Open("./test.txt")
	// 打开文件
	// file=file对象 file=file指针 file=file文件句柄
	if err != nil {
		fmt.Println("file error",err)
	}

	fmt.Printf("file=%v\n",file)

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		// 读到一个换行就结束
		str, err := reader.ReadString('\n')

		// 读取结束， 结束死循环
		if err == io.EOF {
			break
		}

		fmt.Printf(str)
	}

	fmt.Println("文件读取结束...")


	// ioutil.ReadFile一次性读取文件
	// 不需要显式的open文件，close文件
	file1 := "./test1.txt"
	content, err := ioutil.ReadFile(file1)

	if err != nil {
		fmt.Printf("read file err = %v\n",err)
	}

	fmt.Printf("%v",string(content))
}

















