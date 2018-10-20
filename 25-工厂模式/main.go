

package main 

import (
	"fmt"
)

type student struct {
	Name string
	score float64
}

func NewStudent(s string) *student {
	return &student {
		Name: s,
	}
}


// 如果字段首字母小写，则在其他包不可以直接使用，提供一个方法供其使用
func (s *student) GetScore() float64 {
	return s.score
}

// go的结构体没有构造函数，通常可以使用工厂模式解决这个问题
// 使用工厂模式实现挎包创建结构体变量（实例）
func main(){
	var s = NewStudent("haha")
	fmt.Println(*s)
	fmt.Println("name=",s.Name)
	fmt.Println(s.GetScore())
}