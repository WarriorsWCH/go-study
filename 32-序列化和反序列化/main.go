

package main 

import (
	"fmt"
	"encoding/json"
)

// json序列化
// json序列化是指，将有key-value结构的数据类型（比如结构体、map、切片）序列化成json字符串的操作
// 这里我们介绍一下结构体、map和切片的序列化，其他数据类型的序列化类似

// 对于结构体的序列化，如果我们希望小序列化后的key的名字，又我们自己重新制定，那么可以给strcut指定一个tag标签

type Monster struct {
	Name string `json:"monster_name"`
	Age int
	Birthday string
	Sal string
	Skill string
} 

func main(){

	monster := Monster {
		Name: "八戒",
		Birthday: "1928-3-21",
		Age: 20,
		Sal: "1000",
		Skill: "eating",
	}

	data, err := json.Marshal(&monster)
	// 反序列
	// json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	fmt.Printf("序列化后=%v\n",string(data))
	
	testMap()

	testSlice()
}


func testMap() {
	// key是string value为任意类型
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "haha"
	a["age"] = 20

	// map本身就是引用类型
	data, err := json.Marshal(a)
	// 反序列化
	// 注意：反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	// err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}
	fmt.Printf("序列化后=%v\n",string(data))
}


func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 33
	slice = append(slice, m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 44
	slice = append(slice,m2)

	data, err := json.Marshal(slice)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n",err)
	}

	fmt.Printf("序列化后=%v\n",string(data))
}

// 在序列化一个json字符串时，要确保反序列化后的数据类型和原来序列化前的数据类型一致
// 如果json字符串是通过程序获取的，则不需要对"转义处理












