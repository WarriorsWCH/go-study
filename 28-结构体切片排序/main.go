

package main 

import (
	"fmt"
	"sort"
	"math/rand"
)
// 结构体切片排序
type Hero struct {
	Name string
	Age int
}

// 声明一个Hero结构体切片类型
type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main(){

	var intSlice = []int{0,-1,7,20,44}
	sort.Ints(intSlice)
	fmt.Println(intSlice)//[-1 0 7 20 44]

	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄|%d",rand.Intn(10)),
			Age:rand.Intn(10),
		}
		heroes = append(heroes,hero)
	}

	for _, v := range heroes{
		fmt.Println(v)
	}

	sort.Sort(heroes)
	// 要使用sort.Sort得实现Len()/Less()/Swap()接口
	fmt.Println("排序后")
	for _, v := range heroes {
		fmt.Println(v)
	}

}















