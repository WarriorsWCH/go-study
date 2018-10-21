

package main 

import(
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"io"
)

func main(){
	// 创建一个新写文件，写入5句话
	filePath := "./f1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	// const (
      //   O_RDONLY int = syscall.O_RDONLY // open the file read-only.
      //   O_WRONLY int = syscall.O_WRONLY // open the file write-only.
      //   O_RDWR int = syscall.O_RDWR // open the file read-write.
      //   O_APPEND int = syscall.O_APPEND // append data to the file when writing.
      //   O_CREATE int = syscall.O_CREAT // create a new file if none exists.
      //   O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist
      //   O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
      //   O_TRUNC int = syscall.O_TRUNC // if possible, truncate file when opened.
     // )
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	defer file.Close()

	str := "hello world\n"
	// 写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer是带缓存的，因此在调用writer的时候，其实内容先是写到缓存中，
	// 所以需要使用Flush方法，将缓存的数据真正的写入到文件中
	writer.Flush()

	overWrite()
	// writeCopy()

	fmt.Println(PathExists("./f1.txt"))
	fmt.Println(PathExists("./f4.txt"))

	w, e := CopyFile("./f3.txt", "./f1.txt")
	if e != nil {
		panic(err)
	}
	fmt.Println(w, e)

	Count()
}

// 覆盖写入，追加写入
func overWrite() {
	f1Path := "./f1.txt"
	file, err := os.OpenFile(f1Path, os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	str := "你好\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}



// 将一个文件内容写入另一个文件
func writeCopy() {
	file1Path := "./f1.txt"
	file2Path := "./f2.txt"

	data, err := ioutil.ReadFile(file1Path)
	fmt.Println(string(data))
	if err != nil {
		panic(err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		panic(err)
	}
}

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	fileinfo, err := os.Stat(path)
	// 文件存在
	if err == nil {
		fmt.Println(fileinfo.Name())    //获取文件名
	    fmt.Println(fileinfo.IsDir())   //判断是否是目录，返回bool类型
	    fmt.Println(fileinfo.ModTime()) //获取文件修改时间
	    fmt.Println(fileinfo.Mode())
	    fmt.Println(fileinfo.Size()) //获取文件大小
	    fmt.Println(fileinfo.Sys())
		return true, nil
	}

	// 文件或文件夹存在
	if os.IsNotExist(err) {
		return false, nil
	}
	// 如果错误类型是其他类型，不去定是否存在
	return false, err
}


func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
	}

	defer srcFile.Close()
	// 通过srcFile获取到Reader
	reader := bufio.NewReader(srcFile)

	// 打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	// 通过dstFile获取到writer
	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writer, reader)
}

// 统计英文、数字。空格和其他字符常量
type CharCount struct {
	ChCount int //英文个数
	NumCount int //记录数字个数
	SpaceCount int //空格个数
	OtherCount int //其他个数
}

func Count() {
	fileName := "./f3.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
	}

	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		newStr := []rune(str)
		for _, v := range newStr {
			switch {
				case v >= 'a' && v <= 'z':
					fallthrough
				case v >= 'A' && v <= 'Z':
					count.ChCount++
				case v == ' ' || v == '\t':
					count.SpaceCount++
				case v >= '0' && v <= '9':
					count.NumCount++
				default:
					count.OtherCount++
			}
		}

		fmt.Printf("字符的个数为：%v, 数字的个数为：%v，空格的个数为：%v,其他个数为：%v\n",
			count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
	}
}







































