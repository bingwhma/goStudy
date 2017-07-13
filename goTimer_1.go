package main

import (
	"fmt"
	"time"
)

var count int = 0

func doSomething() int {
	count++
	return count
}

func f1() {
	fmt.Println("f1 done !")
}

func f2() {
	fmt.Println("f2 done !")
}

func main() {

	t := time.Tick(2 * time.Second)
	fmt.Println(time.Now())

	i := 0
	for now := range t {
		fmt.Println(now, doSomething())
		i++
		if i > 10 {
			break
		}
	}

	time.AfterFunc(5*time.Second, f1)
	time.AfterFunc(2*time.Second, f2)
	fmt.Println("main thread")
	time.Sleep(10 * time.Second)
	fmt.Println("=====over========")
}
