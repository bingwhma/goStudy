package main

import (
	"fmt"
	"sync"
	"time"
)

var lock *sync.Mutex

func lockPrint(i int) {
	fmt.Println(time.Now(), i, "lock start")
	lock.Lock()

	fmt.Println(time.Now(), i, " in lock")
	time.Sleep(4 * time.Second)
	lock.Unlock()
	fmt.Println(time.Now(), i, " unlock")
}

func main() {
	lock = new(sync.Mutex)

	go lockPrint(1)
	lockPrint(2)

	time.Sleep(4 * time.Second)
	fmt.Println("=====over======")
}
