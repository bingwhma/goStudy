package main

import (
	"fmt"
)

type Person struct {
	ID      string
	Name    string
	Address string
}

func main() {
	personDB := make(map[string](Person))

	var person = Person{"12346", "Xym", "bbb"}
	personDB["p1"] = person

	p, ok := personDB["p1"]
	if !ok {
		fmt.Println("no data found")
	}

	//打印出全部值 和各个分值
	fmt.Printf(" \n personDB[p1] = \n {ID=%v \n Name=%v \n address=%v \n}", p.ID, p.Name, p.Address)

	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "cc"
	m1["d"] = "dd"
	fmt.Println("\n", m1) //输出元素的顺序是随机的，go语言中map是无序的

	// 直接创建
	m2 := make(map[string]string)
	m2["a"] = "aa"
	m2["b"] = "bb"
	fmt.Println("\n", m2)

	// 初始化 + 赋值一体化 注意最后一个元素后面也需要带上逗号的
	m3 := map[string]string{
		"a":  "aa",
		"b":  "bb",
		"c":  "cc",
		"xx": "xx",
		"d":  "dd", //注意最后一个元素后面也需要带上逗号的
	}
	fmt.Println("\n", m3)

	//删除一个key对应的元素 map的删除操作
	fmt.Println("删除之前", m3)
	delete(m3, "a")
	fmt.Println("删除之后", m3)

	fmt.Println("output map")
	for k, v := range m3 {
		fmt.Printf("key=%v, val=%v \n", k, v)
	}
}
