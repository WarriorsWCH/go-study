

package main 

import (
	"fmt"
)


type Monkey struct {
	Name string
}

type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

type LittleMonkey struct {
	Monkey //继承
	// Name string //接口
}

func (this Monkey) climbling(){
	fmt.Println(this.Name, "生来会爬树")
}

func (this LittleMonkey) Swimming() {
	fmt.Println(this.Name,"通过学习，我会游泳了")
}
func (this LittleMonkey) Flying() {
	fmt.Println(this.Name, "通过学习，我会飞翔了")
}

func main() {

	testInherit()

	// testInterface()
}


func testInherit() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbling()
	// monkey.Flying()
	monkey.Swimming()

	bigmonkey := Monkey{
		Name: "六耳猕猴",
	}

	monkey.Monkey = bigmonkey
	bigmonkey.climbling()
	monkey.climbling()
	monkey.Swimming()
	fmt.Println(monkey.Monkey.Name,monkey.Name)
}

func testInterface(){
	// var monkey  = LittleMonkey{"猴子"}
	// var b BirdAble = monkey
	// b.Flying()

	// var f FishAble = monkey
	// f.Swimming()
}

// 当A结构体继承了B结构体，nameA结构体就自动继承了B结构体的字段和方法，并且可以直接使用
// 当A结构体需要扩展功能，同时不希望去破坏继承关系，则可以去实现某个接口即可，
// 因此我们可以认为:实现接口是对继承的补充

// 接口比继承更加灵活，继承是满足is - a的关系，而接口只需要满足like - a的关系
// 接口在一定程度上实现代码的解耦














