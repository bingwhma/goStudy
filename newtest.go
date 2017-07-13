package main

import "fmt"

type Student struct {
	id   int
	name string
}

func main() {
	// 这是一个用来分配内存的内置函数，它的第一个参数是一个类型，不是一个值，
	// 它的返回值是一个指向新分配的 t 类型的零值的指针
	var s_1 *Student = new(Student)
	s_1.id = 100
	s_1.name = "cat"

	// 直接使用struct{} 来初始化strut时，返回的是一个struct类型的值，
	// 而不是指针两者是不一样的
	var s_2 Student = Student{200, "xxxx"}

	// 从上面代码的声明和打印的结果中就可以看出 s_1 的类型为指针，s_2 为一个Student类型
	fmt.Println(s_1, s_2)

	fmt.Println(s_1, &s_2)
}
