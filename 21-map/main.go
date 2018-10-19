

package main 


import(
	"fmt"
	"sort"
)

func main(){
	// map
	// map是key-value数据结构，又称为字段或者关联数组，类似其他语言的集合
	// key可以是什么类型：
	// go中的map的key，可以是很多种类型，比如nool，string，指针，channel
	// 还可以是只包含前面几个类型的接口，结构体，数组，通常key为int，string
	// 注意：slice，map还有function不可以，因为这几个没法用==来判断
	// value可以的类型和key基本一样

	var a map[string]string
	// 注意：声明是不会分配内存的，初始化需要make分配内存后才能赋值和使用
	a = make(map[string]string,10)
	a["no1"] = "hello"
	a["no2"] = "world"
	fmt.Println(a)

	// 细节：
	// 1.map在使用前一定要make
	// 2.map的key是不能重复的，如果重复了，则以最后这个key-value为准
	// 3.map的value是可以相同的
	// 4.map的key-value是无序的
	// 5.make内置函数数目


	// 第二种方式
	cities := make(map[string]string)
	cities["add1"] = "北京"
	cities["add2"] = "上海"
	fmt.Println(cities)

	// 第三种方式
	heros := map[string]string{
		"hero1": "超人",
		"hero2": "蜘蛛侠",
	}
	fmt.Println(heros)


	stus := make(map[string]map[string]string)
	stus["stu1"] = make(map[string]string,3)
	stus["stu1"]["name"] = "tom"
	stus["stu1"]["sex"] = "男"
	stus["stu1"]["age"] = "18"

	stus["stu2"] = make(map[string]string,3)
	stus["stu2"]["name"] = "jack"
	stus["stu2"]["sex"] = "男"
	stus["stu2"]["age"] = "20"

	fmt.Println(stus)

	// map增加和更新
	// map["key"] = value key存在就是修改，key不存在就是增加
	stus["stu1"]["name"] = "刘老怪"
	fmt.Println(stus)

	// delete是一个内置函数，如果key存在，就删除，key不存在，就不操作，也不会报错
	delete(stus,"stu2")
	fmt.Println(stus)

	val, ok := cities["add2"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("不存在")
	}

	// map排序
	// 1.ho中没有一个专门的方法针对map的key进行排序
	// 2.go中的map默认是无序的，注意也不是按照添加顺序存放的，你每次遍历，得到的输出可能不一样
	// 3.go中的map的排序，是先将key进行排序，然后根据key值遍历输出即可
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 30
	map1[4] = 44
	fmt.Println(map1)

	var keys []int
	for _, k := range map1 {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	users := make(map[string]map[string]string,10)
	users["smith"] = make(map[string]string,2)
	users["smith"]["pwa"] = "99999"
	users["smith"]["nick"] = "love"

	// 细节说明
	// 1.map是引用类型，遵守引用类型传递的机制，在一个函数接受map，修改后庙会直接修改原来的map
	// 2.map的容量达到后，再像map增加元素，会自动扩容，并不会发生panic，也就是说map能动态的增加键值对
	// 3，map的value也经常使用struct类型，更适合管理复杂的数据，比value是一个map更好

	modifyUser(users,"tom")
	fmt.Println(users)
	modifyUser(users,"haha")
	fmt.Println(users)
	modifyUser(users,"smith")
	fmt.Println(users)

}


func modifyUser(users map[string]map[string]string,name string) {
	if users[name] != nil {
		users[name]["pwa"] = "888888"
	} else {
		users[name] = make(map[string]string,2)
		users[name]["pwa"] = "66666"
		users[name]["nick"] = "ok"
	}
}
























